package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/Zmey56/go-kit-user-service/internal/service"
)

// Request and Response struct for endpoints
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type CreateUserResponse struct {
	User  service.User `json:"user"`
	Error string       `json:"error,omitempty"`
}

type GetUserRequest struct {
	ID int `json:"id"`
}

type GetUserResponse struct {
	User  service.User `json:"user"`
	Error string       `json:"error,omitempty"`
}

type UpdateUserRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UpdateUserResponse struct {
	User  service.User `json:"user"`
	Error string       `json:"error,omitempty"`
}

type DeleteUserRequest struct {
	ID int `json:"id"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// Endpoints contains all the microservice endpoints
type Endpoints struct {
	CreateUserEndpoint endpoint.Endpoint
	GetUserEndpoint    endpoint.Endpoint
	UpdateUserEndpoint endpoint.Endpoint
	DeleteUserEndpoint endpoint.Endpoint
}

// Endpoints
func MakeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		user, err := svc.CreateUser(req.Name, req.Email, req.Age)
		if err != nil {
			return CreateUserResponse{Error: err.Error()}, nil
		}
		return CreateUserResponse{User: user}, nil
	}
}

func MakeGetUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := svc.GetUserByID(req.ID)
		if err != nil {
			return GetUserResponse{Error: err.Error()}, nil
		}
		return GetUserResponse{User: user}, nil
	}
}

func MakeUpdateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		user, err := svc.UpdateUser(req.ID, req.Name, req.Email, req.Age)
		if err != nil {
			return UpdateUserResponse{Error: err.Error()}, nil
		}
		return UpdateUserResponse{User: user}, nil
	}
}

func MakeDeleteUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		err := svc.DeleteUser(req.ID)
		if err != nil {
			return DeleteUserResponse{Error: err.Error()}, nil
		}
		return DeleteUserResponse{Message: "User deleted successfully"}, nil
	}
}
