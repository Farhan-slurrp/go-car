package user

import (
	"context"
	"errors"

	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
)

type userService struct{}

func NewService() Service { return &userService{} }

func (w *userService) GetUserData(_ context.Context, id string) (*internal.User, error) {
	user := internal.User{}
	if id == "" {
		return nil, errors.New("user not found")
	}
	database.DB.Find(&user, "ID = ?", id)
	return &user, nil
}

func (w *userService) UpdateUserData(_ context.Context, id string, user *internal.User) (uint, error) {
	oldUser := *user
	if result := database.DB.Find(&oldUser, "ID = ?", id); result.Error != nil {
		return 0, result.Error
	}
	if user.Name != "" {
		oldUser.Name = user.Name
	}
	if user.Email != "" {
		oldUser.Email = user.Email
	}
	if user.Role != "" {
		oldUser.Role = user.Role
	}
	database.DB.Save(&oldUser)
	return oldUser.ID, nil
}
