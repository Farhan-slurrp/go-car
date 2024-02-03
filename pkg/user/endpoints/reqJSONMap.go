package endpoints

import "github.com/Farhan-slurrp/go-car/internal"

type GetUserDataRequest struct {
	ID string `json:"id"`
}

type GetUserDataResponse struct {
	User *internal.User `json:"user,omitempty"`
	Err  string         `json:"error,omitempty"`
}

type AuthorizeUserTokenRequest struct {
	Token string `json:"token"`
}

type AuthorizeUserTokenResponse struct {
	User *internal.User `json:"user,omitempty"`
	Err  string         `json:"error,omitempty"`
}

type UpdateUserDataRequest struct {
	ID   string         `json:"id"`
	User *internal.User `json:"user"`
}

type UpdateUserDataResponse struct {
	Id  uint   `json:"id"`
	Err string `json:"error,omitempty"`
}
