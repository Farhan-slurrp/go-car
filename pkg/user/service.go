package user

import (
	"context"

	"github.com/Farhan-slurrp/go-car/internal"
)

type Service interface {
	GetUserData(ctx context.Context, email string) (internal.User, error)
	UpdateUserData(ctx context.Context, user internal.User) (uint, error)
}
