package endpoints

import "github.com/Farhan-slurrp/go-car/internal"

type LoginRequest struct{}

type LoginResponse struct {
	Status string `json:"status"`
	Err    string `json:"err,omitempty"`
}

type CallbackRequest struct{}

type CallbackResponse struct {
	Status string `json:"token"`
	Err    string `json:"err,omitempty"`
}

type GetUserDataRequest struct {
	UserId string `json:"user_id"`
}

type GetUserDataResponse struct {
	User internal.User `json:"user"`
	Err  string        `json:"err,omitempty"`
}
