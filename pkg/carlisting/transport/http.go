package transport

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
	"github.com/Farhan-slurrp/go-car/pkg/carlisting/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
	userHost = "https://caruserr-farhannurzi.koyeb.app/"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	database.ConnectDB()
	database.DB.AutoMigrate(&internal.CarListing{})

	m := mux.NewRouter()

	m.Methods("GET", "POST").Path("/cars").Handler(httptransport.NewServer(
		ep.GetCarListingsEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))

	m.Methods("POST").Path("/cars/create").Handler(httptransport.NewServer(
		ep.CreateListingEndpoint,
		decodeHTTPCreateListingRequest,
		encodeResponse,
	))

	m.Methods("PATCH", "PUT").Path("/cars/{id}/update").Handler(httptransport.NewServer(
		ep.UpdateListingEndpoint,
		decodeHTTPUpdateListingRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetCarListingsRequest
	if r.ContentLength == 0 {
		return req, nil
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPCreateListingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return nil, errors.New("authorization required")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	user := internal.User{}
	authorizeResp := internal.AuthorizeResponse{}
	resp, userErr := http.Get(userHost + reqToken)
	if userErr != nil {
		return nil, userErr
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("authorization required")
	}
	err = json.Unmarshal(data, &authorizeResp)
	if err != nil {
		return nil, errors.New("please log in")
	}

	user = *authorizeResp.User
	if user.Role != "host" {
		return nil, errors.New("method not allowed")
	}

	var req endpoints.CreateListingRequest
	reqErr := json.NewDecoder(r.Body).Decode(&req)
	if reqErr != nil {
		return nil, reqErr
	}
	req.CarListing.OwnerId = user.ID
	return req, nil
}

func decodeHTTPUpdateListingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return nil, errors.New("authorization required")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	user := internal.User{}
	authorizeResp := internal.AuthorizeResponse{}
	resp, userErr := http.Get(userHost + reqToken)
	if userErr != nil {
		return nil, userErr
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("authorization required")
	}
	authErr := json.Unmarshal(data, &authorizeResp)
	if authErr != nil {
		return nil, errors.New("please log in")
	}

	user = *authorizeResp.User
	if user.Role != "host" {
		return nil, errors.New("method not allowed")
	}

	var req endpoints.UpdateListingRequest
	req.ID = id
	err = json.NewDecoder(r.Body).Decode(&req)
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

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
