package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Farhan-slurrp/go-car/authentication"
	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
	"github.com/Farhan-slurrp/go-car/internal/config"
	"github.com/Farhan-slurrp/go-car/pkg/user/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	database.ConnectDB()
	database.DB.AutoMigrate(&internal.User{})
	config.GoogleConfig()

	m := mux.NewRouter()

	m.HandleFunc("/login", handleLogin)
	m.HandleFunc("/callback", handleCallback)

	m.Methods("GET").Path("/users/{id}").Handler(httptransport.NewServer(
		ep.GetUserDataEndpoint,
		decodeHTTPGetUserDataRequest,
		encodeResponse,
	))

	m.Methods("GET").Path("/authorize/{token}").Handler(httptransport.NewServer(
		ep.AuthorizeUserTokenEndpoint,
		decodeHTTPAuthorizeUserTokenRequest,
		encodeResponse,
	))

	m.Methods("PATCH", "PUT").Path("/users/{id}/update").Handler(httptransport.NewServer(
		ep.UpdateUserDataEndpoint,
		decodeHTTPUpdateUserDataRequest,
		encodeResponse,
	))

	return m
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := authentication.GenerateStateOauthCookie(w)
	u := config.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := authentication.GetToken(code)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data, err := authentication.GetUserDataFromGoogle(token.AccessToken)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user := internal.User{}
	userResp := internal.UserResponse{}
	err = json.Unmarshal(data, &userResp)
	if err != nil {
		fmt.Printf("%v", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user.Email = userResp.Email
	user.Name = userResp.Name
	user.Role = "user"
	database.DB.Create(&user)

	fmt.Fprintf(w, "Access token: %s\n", token.AccessToken)
}

func decodeHTTPGetUserDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	return endpoints.GetUserDataRequest{ID: id}, nil
}

func decodeHTTPAuthorizeUserTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	token, ok := vars["token"]
	if !ok {
		return nil, ErrBadRouting
	}

	return endpoints.AuthorizeUserTokenRequest{Token: token}, nil
}

func decodeHTTPUpdateUserDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	var req endpoints.UpdateUserDataRequest
	req.ID = id
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	// case util.ErrUnknown:
	// 	w.WriteHeader(http.StatusNotFound)
	// case util.ErrInvalidArgument:
	// 	w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
