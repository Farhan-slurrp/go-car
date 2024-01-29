package endpoints

import (
	"context"
	"errors"

	"github.com/Farhan-slurrp/go-car/pkg/carlisting"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetEndpoint             endpoint.Endpoint
	ListNewCarEndpoint      endpoint.Endpoint
	SetAvailabilityEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc carlisting.Service) Set {
	return Set{
		GetEndpoint:             MakeGetEndpoint(svc),
		ListNewCarEndpoint:      MakeListNewCarEndpoint(svc),
		SetAvailabilityEndpoint: MakeSetAvailabilityEndpoint(svc),
	}
}

func MakeGetEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		docs, err := svc.Get(ctx)
		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}
		return GetResponse{docs, ""}, nil
	}
}

func MakeListNewCarEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListNewCarRequest)
		Int, err := svc.ListNewCar(ctx, req.B)
		if err != nil {
			return GetResponse{Int, err.Error()}, nil
		}
		return GetResponse{Int, ""}, nil
	}
}

func MakeSetAvailabilityEndpoint(svc carlisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetAvailabilityRequest)
		Int, err := svc.SetAvailability(ctx, req.Dates)
		if err != nil {
			return GetResponse{Int, err.Error()}, nil
		}
		return GetResponse{Int, ""}, nil
	}
}

func (s *Set) Get(ctx context.Context) (int, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest)
	if err != nil {
		return 0, err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return 0, errors.New(getResp.Err)
	}
	return getResp.Int, nil
}

func (s *Set) ListNewCar(ctx context.Context) (int, error) {
	resp, err := s.ListNewCarEndpoint(ctx, ListNewCarRequest{B: int})
	if err != nil {
		return 0, err
	}
	listNewCarResp := resp.(ListNewCarResponse)
	if listNewCarResp.Err != "" {
		return 0, errors.New(listNewCarResp.Err)
	}
	return listNewCarResp.Int, nil
}

func (s *Set) SetAvailability(ctx context.Context) (int, error) {
	resp, err := s.SetAvailabilityEndpoint(ctx, SetAvailabilityRequest{Dates: string})
	if err != nil {
		return 0, err
	}
	setAvailabilityResp := resp.(SetAvailabilityResponse)
	if setAvailabilityResp.Err != "" {
		return 0, errors.New(setAvailabilityResp.Err)
	}
	return setAvailabilityResp.Int, nil
}
