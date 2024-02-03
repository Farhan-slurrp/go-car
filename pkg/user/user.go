package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Farhan-slurrp/go-car/authentication"
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

func (w *userService) AuthorizeUserToken(_ context.Context, token string) (*internal.User, error) {
	userData, err := authentication.GetUserDataFromGoogle(token)
	if err != nil {
		return nil, errors.New("authorization failed")
	}
	user := internal.User{}
	userResp := internal.UserResponse{}
	userErr := json.Unmarshal(userData, &userResp)
	if userErr != nil {
		return nil, userErr
	}

	database.DB.Find(&user, "email = ?", userResp.Email)
	if user.Email == "" {
		return nil, errors.New("user not found")
	}

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
