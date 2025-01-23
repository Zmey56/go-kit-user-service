package main

import (
	"net/http"

	"github.com/Zmey56/go-kit-user-service/internal/endpoint"
	"github.com/Zmey56/go-kit-user-service/internal/service"
)

func main() {
	// Initializing the service
	svc := service.NewUserService()

	// Creating endpoints
	createUserEndpoint := endpoint.MakeCreateUserEndpoint(svc)
	getUserEndpoint := endpoint.MakeGetUserEndpoint(svc)
	updateUserEndpoint := endpoint.MakeUpdateUserEndpoint(svc)
	deleteUserEndpoint := endpoint.MakeDeleteUserEndpoint(svc)

	// Using endpoints in the transport layer
	_ = createUserEndpoint
	_ = getUserEndpoint
	_ = updateUserEndpoint
	_ = deleteUserEndpoint

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
