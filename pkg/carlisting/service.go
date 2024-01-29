package carlisting

import (
	"context"
)

type Service interface {
	// Get the list of all documents
	Get(ctx context.Context) (int, error)
	ListNewCar(ctx context.Context, b int) (int, error)
	SetAvailability(ctx context.Context, Dates string) (int, error)
}
