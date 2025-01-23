package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	kitHttp "github.com/go-kit/kit/transport/http"

	"github.com/Zmey56/go-kit-user-service/internal/endpoint"
)

// NewHTTPHandler creates a new HTTP handlers for all endpoints
func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/users", kitHttp.NewServer(
		endpoints.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
	))

	mux.Handle("/users/", kitHttp.NewServer(
		endpoints.GetUserEndpoint,
		decodeGetUserRequest,
		encodeResponse,
	))

	mux.Handle("/users/update", kitHttp.NewServer(
		endpoints.UpdateUserEndpoint,
		decodeUpdateUserRequest,
		encodeResponse,
	))

	mux.Handle("/users/delete", kitHttp.NewServer(
		endpoints.DeleteUserEndpoint,
		decodeDeleteUserRequest,
		encodeResponse,
	))

	return mux
}

// Query decoding functions

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return nil, http.ErrMissingFile
	}
	return endpoint.GetUserRequest{ID: atoi(id)}, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return nil, http.ErrMissingFile
	}
	return endpoint.DeleteUserRequest{ID: atoi(id)}, nil
}

// Response encoding function

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// atoi converts a string to an integer
func atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
