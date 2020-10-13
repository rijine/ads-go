package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgryski/trifles/uuid"
	"github.com/rijine/ads-api/pkg/counters"
	"github.com/rijine/ads-api/pkg/notify"
	"log"
	"strings"
	"time"

	"github.com/rijine/ads-api/internal/config"
	"github.com/rijine/ads-api/internal/database"
	"github.com/rijine/ads-api/pkg/graph/model"
	"github.com/rijine/ads-api/pkg/jwts"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSrv     = jwts.NewJwtService(config.JwtConf)
	userRepo   = NewUserRepository()
	counterSrv = counters.NewCounterService()
	notifySrv  = notify.NewNotifyService()
)

type Service interface {
	Register(userForm *model.NewUser) (bool, error)
	Users() ([]*model.User, error)
	Login(credential *model.Credential) (*model.AuthUser, error)
}

type service struct{}

func NewUserService() Service {
	fmt.Print("user service")
	return &service{}
}

func (s *service) Users() ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := database.Collection("users").Find(ctx, bson.M{})

	if err != nil {
		log.Print(err)
	}

	var ss = make([]*model.User, 0)
	// var ss1 = make([]*User, 0)
	err = cur.All(context.TODO(), &ss)

	// fmt.Print(ss1)
	/*for cur.Next(context.Background()) {
		var s model.User
		var d User
		if err := cur.Decode(&s); err == nil {
			// c.Send(err)
			// ss = append(ss, s)
			fmt.Print(s, d)
		}
	}*/

	return ss, err
}

func (s *service) Register(userForm *model.NewUser) (bool, error) {
	user, err := userRepo.GetUser(userForm.Email)
	if user != nil && user.Status != NotApproved {
		return false, errors.New("user already exists")
	}

	count, err := counterSrv.Count("users")
	if err != nil {
		return false, errors.New("something went wrong, try again")
	}

	url := strings.Join([]string{
		strings.ToLower(userForm.FirstName),
		strings.ToLower(userForm.LastName),
		fmt.Sprint(count),
	}, "-")
	key := uuid.UUIDv4()
	expiry := time.Now().Add(time.Hour * 24).Unix()
	bs, _ := bcrypt.GenerateFromPassword([]byte(userForm.Password), bcrypt.DefaultCost)
	newUser := User{
		FirstName:          userForm.FirstName,
		LastName:           userForm.LastName,
		Username:           userForm.Email,
		Email:              userForm.Email,
		Password:           string(bs),
		Status:             NotApproved,
		Url:                url,
		VerificationKey:    key,
		VerificationExpiry: expiry,
	}

	_, err = userRepo.AddUser(&newUser)
	if err != nil {
		return false, errors.New("failed to register, please try again")
	}

	defer func() {
		err := notifySrv.VerifyEmail("e.rijin@gmail.com")
		if err != nil {
			log.Println("email sending failed, changed")
		}
	}()

	return true, nil
}

func (s *service) Login(credential *model.Credential) (*model.AuthUser, error) {
	user, err := userRepo.GetUser(credential.Username)

	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if user.Status != Approved {
		return nil, errors.New("user is not verified or blocked")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))

	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	token, expiry, err := jwtSrv.Generate(credential.Username)

	if err != nil {
		return nil, errors.New("something went wrong please contact admin")
	}

	authUser := model.AuthUser{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		Token:       token,
		Expiry:      int(expiry),
	}

	return &authUser, nil
}
