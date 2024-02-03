package user

import (
	"context"

	"github.com/Farhan-slurrp/go-car/internal"
)

type Service interface {
	GetUserData(ctx context.Context, id string) (*internal.User, error)
	AuthorizeUserToken(ctx context.Context, token string) (*internal.User, error)
	UpdateUserData(ctx context.Context, id string, user *internal.User) (uint, error)
}
