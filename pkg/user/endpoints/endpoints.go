package endpoints

import (
	"context"
	"errors"

	"github.com/Farhan-slurrp/go-car/internal"
	"github.com/Farhan-slurrp/go-car/pkg/user"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetUserDataEndpoint        endpoint.Endpoint
	AuthorizeUserTokenEndpoint endpoint.Endpoint
	UpdateUserDataEndpoint     endpoint.Endpoint
}

func NewEndpointSet(svc user.Service) Set {
	return Set{
		GetUserDataEndpoint:        MakeGetUserDataEndpoint(svc),
		AuthorizeUserTokenEndpoint: MakeAuthorizeUserTokenEndpoint(svc),
		UpdateUserDataEndpoint:     MakeUpdateUserDataEndpoint(svc),
	}
}

func MakeGetUserDataEndpoint(svc user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserDataRequest)
		user, err := svc.GetUserData(ctx, req.ID)
		if err != nil {
			return GetUserDataResponse{user, err.Error()}, nil
		}
		return GetUserDataResponse{user, ""}, nil
	}
}

func MakeAuthorizeUserTokenEndpoint(svc user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthorizeUserTokenRequest)
		user, err := svc.AuthorizeUserToken(ctx, req.Token)
		if err != nil {
			return AuthorizeUserTokenResponse{user, err.Error()}, nil
		}
		return AuthorizeUserTokenResponse{user, ""}, nil
	}
}

func MakeUpdateUserDataEndpoint(svc user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserDataRequest)
		id, err := svc.UpdateUserData(ctx, req.ID, req.User)
		if err != nil {
			return UpdateUserDataResponse{id, err.Error()}, nil
		}
		return UpdateUserDataResponse{id, ""}, nil
	}
}

func (s *Set) GetUserData(ctx context.Context, id string) (*internal.User, error) {
	resp, err := s.GetUserDataEndpoint(ctx, GetUserDataRequest{ID: id})
	if err != nil {
		return nil, err
	}
	getUserDataResp := resp.(GetUserDataResponse)
	if getUserDataResp.Err != "" {
		return nil, errors.New(getUserDataResp.Err)
	}
	return getUserDataResp.User, nil
}

func (s *Set) AuthorizeUserToken(ctx context.Context, token string) (*internal.User, error) {
	resp, err := s.GetUserDataEndpoint(ctx, AuthorizeUserTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}
	authorizeuserTokenResp := resp.(GetUserDataResponse)
	if authorizeuserTokenResp.Err != "" {
		return nil, errors.New(authorizeuserTokenResp.Err)
	}
	return authorizeuserTokenResp.User, nil
}

func (s *Set) UpdateUserData(ctx context.Context, id string, user *internal.User) (uint, error) {
	resp, err := s.UpdateUserDataEndpoint(ctx, UpdateUserDataRequest{ID: id, User: user})
	if err != nil {
		return 0, err
	}
	updateUserDataResp := resp.(UpdateUserDataResponse)
	if updateUserDataResp.Err != "" {
		return 0, errors.New(updateUserDataResp.Err)
	}
	return updateUserDataResp.Id, nil
}
