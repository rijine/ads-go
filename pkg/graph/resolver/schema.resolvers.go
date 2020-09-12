package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/rijine/ads-api/pkg/uploads"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rijine/ads-api/pkg/graph/generated"
	"github.com/rijine/ads-api/pkg/graph/model"
	"github.com/rijine/ads-api/pkg/users"
)

var (
	uploadService = uploads.NewCounterService()
	// authsrv = auth.Auth
)

func (r *mutationResolver) AddPost(ctx context.Context, post model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditPost(ctx context.Context, post model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ChangePassword(ctx context.Context, password model.ChangePassword) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) VerifyEmail(ctx context.Context, key string) (*model.AuthUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, user *model.NewUser) (bool, error) {
	return usersService.Register(user)
}

func (r *mutationResolver) UploadImage(ctx context.Context, picture graphql.Upload, kind *string) (*string, error) {
	fmt.Println(picture.Filename, *kind)
	if picture.Size >= 1000000 {
		return nil, errors.New("image size too big")
	}

	res, err := uploadService.UploadImage(ctx, picture, kind)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *mutationResolver) UploadImages(ctx context.Context, pictures []*graphql.Upload) (*string, error) {
	for i, _ := range pictures {
		fmt.Println(i)
	}
	v := "uploaded"
	return &v, nil
}

func (r *queryResolver) Login(ctx context.Context, credentials *model.Credential) (*model.AuthUser, error) {
	// panic(fmt.Errorf("not implemented"))
	return usersService.Login(credentials)
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return usersService.Users()
}

func (r *queryResolver) UserByURL(ctx context.Context, url *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context, id *string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PostByURL(ctx context.Context, url *string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context, filters *model.PostFilter) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]*model.Post, error) {
	return make([]*model.Post, 0), nil
}

func (r *userResolver) History(ctx context.Context, obj *model.User) ([]*model.History, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) ConfirmEmail(ctx context.Context, key *string) (*model.AuthUser, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) ID(ctx context.Context, obj *users.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) ProfileImage(ctx context.Context, obj *users.User) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Address(ctx context.Context, obj *users.User) (*model.Address, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Rating(ctx context.Context, obj *users.User) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Customers(ctx context.Context, obj *users.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

var (
	usersService = users.NewUserService()
)
