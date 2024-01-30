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
	B int `json:"B"`
}

type CreateListingResponse struct {
	Int int    `json:"Int"`
	Err string `json:"err,omitempty"`
}

type UpdateListingRequest struct {
	Dates string `json:"dates"`
}

type UpdateListingResponse struct {
	Int int    `json:"Int"`
	Err string `json:"err,omitempty"`
}
