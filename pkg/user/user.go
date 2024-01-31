package user

import (
	"context"

	"github.com/Farhan-slurrp/go-car/internal"
)

type userService struct{}

func NewService() Service { return &userService{} }

func (w *userService) Login(_ context.Context) (string, error) {

	return "", nil
}

func (w *userService) Callback(_ context.Context) (string, error) {

	return "", nil
}
func (w *userService) GetUserData(_ context.Context, id uint) (internal.User, error) {
	user := internal.User{}
	return user, nil
}
