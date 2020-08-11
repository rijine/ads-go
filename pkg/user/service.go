package user

import "fmt"

type Service interface {
}

type service struct{}

func NewUserService() Service {
	fmt.Print("user service")
	return &service{}
}
