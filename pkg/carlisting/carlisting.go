package carlisting

import (
	"context"
	"fmt"

	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
)

type carListingService struct{}

func NewService() Service { return &carListingService{} }

func (w *carListingService) Get(_ context.Context, filters ...internal.Filter) ([]internal.CarListing, error) {

	cars := []internal.CarListing{}
	database.DB.Find(&cars)
	return cars, nil
}

func (w *carListingService) CreateListing(_ context.Context, carListing *internal.CarListing) (uint, error) {
	newCarListing := carListing
	database.DB.Create(&newCarListing)
	return newCarListing.ID, nil
}

func (w *carListingService) UpdateListing(_ context.Context, carListing *internal.CarListing) (string, error) {
	oldCarListing := *carListing
	if result := database.DB.First(&oldCarListing); result.Error != nil {
		fmt.Printf("%v", result)
		return "", result.Error
	}
	oldCarListing = *carListing
	database.DB.Save(&oldCarListing)
	return "Data updated", nil
}
