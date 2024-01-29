package carlisting

import (
	"context"
)

type carListingService struct{}

func NewService() Service { return &carListingService{} }

func (w *carListingService) Get(_ context.Context) (int, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	return 1, nil
}

func (w *carListingService) ListNewCar(_ context.Context, b int) (int, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	return b, nil
}

func (w *carListingService) SetAvailability(_ context.Context, Dates string) (int, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	return 1, nil
}
