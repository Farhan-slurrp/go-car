package endpoints

import (
	"context"
	"errors"

	"github.com/Farhan-slurrp/go-car/internal"
	"github.com/Farhan-slurrp/go-car/pkg/carlisting"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	CreateListingEndpoint endpoint.Endpoint
	UpdateListingEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc carlisting.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		CreateListingEndpoint: MakeCreateListingEndpoint(svc),
		UpdateListingEndpoint: MakeUpdateListingEndpoint(svc),
	}
}

func MakeGetEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		docs, err := svc.Get(ctx, req.Filters...)
		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}
		return GetResponse{docs, ""}, nil
	}
}

func MakeCreateListingEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateListingRequest)
		Int, err := svc.CreateListing(ctx, req.CarListing)
		if err != nil {
			return CreateListingResponse{Int, err.Error()}, nil
		}
		return CreateListingResponse{Int, ""}, nil
	}
}

func MakeUpdateListingEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateListingRequest)
		Int, err := svc.UpdateListing(ctx, req.CarListing)
		if err != nil {
			return UpdateListingResponse{Int, err.Error()}, nil
		}
		return UpdateListingResponse{Int, ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...internal.Filter) ([]internal.CarListing, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []internal.CarListing{}, err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []internal.CarListing{}, errors.New(getResp.Err)
	}
	return getResp.Cars, nil
}

func (s *Set) CreateListing(ctx context.Context, carListing *internal.CarListing) (uint, error) {
	resp, err := s.CreateListingEndpoint(ctx, CreateListingRequest{CarListing: carListing})
	if err != nil {
		return 0, err
	}
	listNewCarResp := resp.(CreateListingResponse)
	if listNewCarResp.Err != "" {
		return 0, errors.New(listNewCarResp.Err)
	}
	return listNewCarResp.ID, nil
}

func (s *Set) UpdateListing(ctx context.Context, carListing *internal.CarListing) (string, error) {
	resp, err := s.UpdateListingEndpoint(ctx, UpdateListingRequest{CarListing: carListing})
	if err != nil {
		return "", err
	}
	setAvailabilityResp := resp.(UpdateListingResponse)
	if setAvailabilityResp.Err != "" {
		return "", errors.New(setAvailabilityResp.Err)
	}
	return setAvailabilityResp.Message, nil
}
