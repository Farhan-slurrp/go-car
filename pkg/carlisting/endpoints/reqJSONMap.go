package endpoints

import "github.com/Farhan-slurrp/go-car/internal"

type GetRequest struct {
	Filters []internal.Filter `json:"filters"`
}

type GetResponse struct {
	Cars []internal.CarListing `json:"cars"`
	Err  string                `json:"err,omitempty"`
}

type CreateListingRequest struct {
	CarListing *internal.CarListing `json:"new_car_listing"`
}

type CreateListingResponse struct {
	ID  uint   `json:"id"`
	Err string `json:"err,omitempty"`
}

type UpdateListingRequest struct {
	CarListing *internal.CarListing `json:"new_car_listing"`
}

type UpdateListingResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}
