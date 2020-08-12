package users

import (
	"fmt"
	"github.com/rijine/ads-api/pkg/graph/model"
)

const (
	COLLECTION = "users"
)

type Repository interface {
	Login(credentials *model.Credential) model.AuthUser
}

type repository struct{}

func NewUserRepository() Repository {
	fmt.Print("Post Repo")
	return &repository{}
}

func (r *repository) Login(credentials *model.Credential) model.AuthUser {
	return model.AuthUser{}
}
