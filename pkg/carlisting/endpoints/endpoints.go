package endpoints

import (
	"context"
	"errors"

	"github.com/Farhan-slurrp/go-car/internal"
	"github.com/Farhan-slurrp/go-car/pkg/carlisting"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetCarListingsEndpoint endpoint.Endpoint
	CreateListingEndpoint  endpoint.Endpoint
	UpdateListingEndpoint  endpoint.Endpoint
}

func NewEndpointSet(svc carlisting.Service) Set {
	return Set{
		GetCarListingsEndpoint: MakeGetCarListingsEndpoint(svc),
		CreateListingEndpoint:  MakeCreateListingEndpoint(svc),
		UpdateListingEndpoint:  MakeUpdateListingEndpoint(svc),
	}
}

func MakeGetCarListingsEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCarListingsRequest)
		docs, err := svc.GetCarListings(ctx, req.Filters...)
		if err != nil {
			return GetCarListingsResponse{docs, err.Error()}, nil
		}
		return GetCarListingsResponse{docs, ""}, nil
	}
}

func MakeCreateListingEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateListingRequest)
		id, err := svc.CreateListing(ctx, req.CarListing)
		if err != nil {
			return CreateListingResponse{id, err.Error()}, nil
		}
		return CreateListingResponse{id, ""}, nil
	}
}

func MakeUpdateListingEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateListingRequest)
		message, err := svc.UpdateListing(ctx, req.ID, req.CarListing)
		if err != nil {
			return UpdateListingResponse{message, err.Error()}, nil
		}
		return UpdateListingResponse{message, ""}, nil
	}
}

func (s *Set) GetCarListings(ctx context.Context, filters ...internal.Filter) ([]internal.CarListing, error) {
	resp, err := s.GetCarListingsEndpoint(ctx, GetCarListingsRequest{Filters: filters})
	if err != nil {
		return []internal.CarListing{}, err
	}
	getResp := resp.(GetCarListingsResponse)
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

func (s *Set) UpdateListing(ctx context.Context, id string, carListing *internal.CarListing) (string, error) {
	resp, err := s.UpdateListingEndpoint(ctx, UpdateListingRequest{ID: id, CarListing: carListing})
	if err != nil {
		return "", err
	}
	updateListingResp := resp.(UpdateListingResponse)
	if updateListingResp.Err != "" {
		return "", errors.New(updateListingResp.Err)
	}
	return updateListingResp.Message, nil
}
