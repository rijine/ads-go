package posts

import (
	"fmt"
)

const (
	COLLECTION = "posts"
)

type Repository interface {
	// Login(credentials *model.Credential) model.AuthUser
}

type repository struct{}

func NewPostRepository() Repository {
	fmt.Print("Post Repo")
	return &repository{}
}

/*func (r *repository) Login(credentials *model.Credential) model.AuthUser {
	return model.AuthUser{}
}*/
