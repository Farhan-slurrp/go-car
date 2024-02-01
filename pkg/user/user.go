package user

import (
	"context"

	"github.com/Farhan-slurrp/go-car/internal"
)

type userService struct{}

func NewService() Service { return &userService{} }

func (w *userService) GetUserData(_ context.Context, email string) (internal.User, error) {
	user := internal.User{}
	return user, nil
}

func (w *userService) UpdateUserData(_ context.Context, user internal.User) (uint, error) {

	return 0, nil
}
