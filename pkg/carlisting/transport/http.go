package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
	"github.com/Farhan-slurrp/go-car/pkg/carlisting/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	database.ConnectDB()
	database.DB.AutoMigrate(&internal.CarListing{})

	m := http.NewServeMux()

	m.Handle("/cars", httptransport.NewServer(
		ep.GetCarListingsEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))

	m.Handle("/cars/create", httptransport.NewServer(
		ep.CreateListingEndpoint,
		decodeHTTPCreateListingRequest,
		encodeResponse,
	))

	m.Handle("/cars/update", httptransport.NewServer(
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
	var req endpoints.CreateListingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPUpdateListingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateListingRequest
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
