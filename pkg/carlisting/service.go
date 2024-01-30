package carlisting

import (
	"context"

	"github.com/Farhan-slurrp/go-car/internal"
)

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.CarListing, error)
	CreateListing(ctx context.Context, b int) (int, error)
	UpdateListing(ctx context.Context, Dates string) (int, error)
}
