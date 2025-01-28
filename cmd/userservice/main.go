package main

import (
	"log"
	"net/http"
	"os"

	kitLog "github.com/go-kit/kit/log"

	"github.com/Zmey56/go-kit-user-service/internal/endpoint"
	"github.com/Zmey56/go-kit-user-service/internal/middleware"
	"github.com/Zmey56/go-kit-user-service/internal/service"
	"github.com/Zmey56/go-kit-user-service/internal/transport"
)

func main() {
	logger := kitLog.NewLogfmtLogger(os.Stdout)
	svc := service.NewUserService()

	endpoints := endpoint.Endpoints{
		CreateUserEndpoint: middleware.LoggingMiddleware(logger)(endpoint.MakeCreateUserEndpoint(svc)),
		GetUserEndpoint:    middleware.LoggingMiddleware(logger)(endpoint.MakeGetUserEndpoint(svc)),
		UpdateUserEndpoint: middleware.LoggingMiddleware(logger)(endpoint.MakeUpdateUserEndpoint(svc)),
		DeleteUserEndpoint: middleware.LoggingMiddleware(logger)(endpoint.MakeDeleteUserEndpoint(svc)),
	}

	handler := transport.NewHTTPHandler(endpoints)

	log.Fatal(http.ListenAndServe(":8080", handler))

}
