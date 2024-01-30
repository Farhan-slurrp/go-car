package carlisting

import (
	"context"

	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
)

type carListingService struct{}

func NewService() Service { return &carListingService{} }

func (w *carListingService) Get(_ context.Context, filters ...internal.Filter) ([]internal.CarListing, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	cars := []internal.CarListing{}
	database.DB.Find(&internal.CarListing{}).Scan(&cars)
	return cars, nil
}

func (w *carListingService) CreateListing(_ context.Context, b int) (int, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	return b, nil
}

func (w *carListingService) UpdateListing(_ context.Context, Dates string) (int, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	return 1, nil
}
