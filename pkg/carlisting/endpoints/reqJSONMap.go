package endpoints

import "github.com/Farhan-slurrp/go-car/internal"

type GetCarListingsRequest struct {
	Filters []internal.Filter `json:"filters"`
}

type GetCarListingsResponse struct {
	Cars []internal.CarListing `json:"cars"`
	Err  string                `json:"error,omitempty"`
}

type CreateListingRequest struct {
	CarListing *internal.CarListing `json:"new_car_listing"`
}

type CreateListingResponse struct {
	ID  uint   `json:"id"`
	Err string `json:"error,omitempty"`
}

type UpdateListingRequest struct {
	ID         string               `json:"id"`
	CarListing *internal.CarListing `json:"new_car_listing"`
}

type UpdateListingResponse struct {
	Message string `json:"message"`
	Err     string `json:"error,omitempty"`
}
