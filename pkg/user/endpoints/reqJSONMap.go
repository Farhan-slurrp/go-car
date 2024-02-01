package endpoints

import "github.com/Farhan-slurrp/go-car/internal"

type GetUserDataRequest struct {
	Email string `json:"email"`
}

type GetUserDataResponse struct {
	User internal.User `json:"user"`
	Err  string        `json:"err,omitempty"`
}

type UpdateUserDataRequest struct {
	User internal.User `json:"user"`
}

type UpdateUserDataResponse struct {
	Id  uint   `json:"id"`
	Err string `json:"err,omitempty"`
}
