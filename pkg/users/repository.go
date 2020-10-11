package users

import (
	"context"
	"fmt"
	"github.com/rijine/ads-api/internal/database"
	"github.com/rijine/ads-api/pkg/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// TODO: move to constants
const (
	COLLECTION = "users"
)

type Repository interface {
	Login(credential *model.Credential) (*model.AuthUser, error)
	GetUser(id string) (*User, error)
	AddUser(user *User) (*User, error)
}

type repository struct{}

func NewUserRepository() Repository {
	fmt.Print("NewUserRepository")
	return &repository{}
}

// TODO: remove
func (r *repository) Login(credential *model.Credential) (*model.AuthUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := database.Collection(COLLECTION).FindOne(ctx, bson.M{"email": credential.Username})
	var user model.AuthUser
	err := res.Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *repository) GetUser(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := database.Collection(COLLECTION).FindOne(ctx, bson.M{"email": id})
	var user User
	err := res.Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *repository) AddUser(user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.Collection(COLLECTION).InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
