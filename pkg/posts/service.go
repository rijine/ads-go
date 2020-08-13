package posts

import (
	"context"
	"fmt"
	"github.com/rijine/ads-api/internal/database"
	"github.com/rijine/ads-api/pkg/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Service interface {
	// Register(userForm *model.NewUser) (bool, error)
	Posts() ([]*model.Post, error)
}

type service struct{}

func NewUserService() Service {
	fmt.Print("user service")
	return &service{}
}

func (s *service) Posts() ([]*model.Post, error) {
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

	return nil, err
}

/*func (s *service) Register(userForm *model.NewUser) (bool, error) {
	usr := User{
		Username: userForm.Email,
		Email:    userForm.Email,
		Password: userForm.Password,
	}

	// Repo layer
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
*/