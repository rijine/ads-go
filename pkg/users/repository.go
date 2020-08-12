package users

import (
	"fmt"
	"github.com/rijine/ads-api/pkg/graph/model"
)

type Repository interface {
	Login(credentials *model.Credential) model.AuthUser
}

type repo struct{}

func NewUserRepository() Repository {
	fmt.Print("Post Repo")
	return &repo{}
}

func (r *repo) Login(credentials *model.Credential) model.AuthUser {
	return model.AuthUser{}
}
