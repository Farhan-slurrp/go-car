package carlisting

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
	"gorm.io/gorm/clause"
)

type carListingService struct{}

func NewService() Service { return &carListingService{} }

func (w *carListingService) GetCarListings(_ context.Context, filters ...internal.Filter) ([]internal.CarListing, error) {
	cars := []internal.CarListing{}
	validKeys := []string{"daily_price", "car_model", "available_from", "available_to"}
	clauses := make([]clause.Expression, 0)

	for _, filter := range filters {
		if !slices.Contains(validKeys, filter.Key) {
			return cars, errors.New("filter doesnt exist")
		} else {
			switch filter.Key {
			case "daily_price":
				clauses = append(clauses, clause.Eq{Column: "daily_price", Value: filter.Value})
			case "car_model":
				clauses = append(clauses, clause.Eq{Column: "car_model", Value: filter.Value})
			case "available_from":
				clauses = append(clauses, clause.Gte{Column: "available_from", Value: filter.Value})
			case "available_to":
				clauses = append(clauses, clause.Lte{Column: "available_to", Value: filter.Value})
			default:
				return cars, errors.New("filter doesnt exist")
			}
		}
	}
	database.DB.Clauses(clauses...).Find(&cars)
	return cars, nil
}

func (w *carListingService) CreateListing(_ context.Context, carListing *internal.CarListing) (uint, error) {
	if carListing.CarModel == "" {
		return 0, errors.New("data is incorrect")
	}
	newCarListing := carListing
	database.DB.Create(&newCarListing)
	return newCarListing.ID, nil
}

func (w *carListingService) UpdateListing(_ context.Context, id string, carListing *internal.CarListing) (string, error) {
	oldCarListing := *carListing
	if result := database.DB.First(&oldCarListing, "ID = ?", id); result.Error != nil {
		fmt.Printf("%v", result)
		return "", result.Error
	}
	if carListing.CarModel != "" {
		oldCarListing.CarModel = carListing.CarModel
	}
	if carListing.DailyPrice != 0 {
		oldCarListing.DailyPrice = carListing.DailyPrice
	}
	if !carListing.AvailableFrom.IsZero() {
		oldCarListing.AvailableFrom = carListing.AvailableFrom
	}
	if !carListing.AvailableTo.IsZero() {
		oldCarListing.AvailableTo = carListing.AvailableTo
	}
	database.DB.Save(&oldCarListing)
	return "Data updated", nil
}
