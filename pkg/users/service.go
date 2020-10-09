package users

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/rijine/ads-api/internal/config"
	"github.com/rijine/ads-api/internal/database"
	"github.com/rijine/ads-api/pkg/graph/model"
	"github.com/rijine/ads-api/pkg/jwts"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSrv   = jwts.NewJwtService(config.JwtConf)
	userRepo = NewUserRepository()
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

//TODO: Move to Auth?
func (s *service) Register(userForm *model.NewUser) (bool, error) {

	bs, _ := bcrypt.GenerateFromPassword([]byte(userForm.Password), bcrypt.DefaultCost)

	usr := User{
		FirstName: userForm.FirstName,
		LastName:  userForm.LastName,
		Username:  userForm.Email,
		Email:     userForm.Email,
		Password:  string(bs),
	}

	// TODO: Repo layer + validations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := database.Collection("users").InsertOne(ctx, usr)
	if err != nil {
		fmt.Print(err)
		return false, err
	}

	fmt.Print(res)
	return true, nil
}

func (s *service) Login(credential *model.Credential) (*model.AuthUser, error) {
	user, err := userRepo.GetUser(credential.Username)

	if err != nil {
		return nil, errors.New("invalid username or password")
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
