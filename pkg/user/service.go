package user

import (
	"context"

	"github.com/Farhan-slurrp/go-car/internal"
)

type Service interface {
	Login(ctx context.Context) (string, error)
	Callback(ctx context.Context) (string, error)
	GetUserData(ctx context.Context, id uint) (internal.User, error)
}
