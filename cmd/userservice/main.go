package main

import (
	"database/sql"
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
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	logger := kitLog.NewLogfmtLogger(os.Stdout)
	svc := service.NewUserService(db)

	endpoints := endpoint.Endpoints{
		CreateUserEndpoint: middleware.LoggingMiddleware(logger)(endpoint.MakeCreateUserEndpoint(svc)),
		GetUserEndpoint:    middleware.LoggingMiddleware(logger)(endpoint.MakeGetUserEndpoint(svc)),
		UpdateUserEndpoint: middleware.LoggingMiddleware(logger)(endpoint.MakeUpdateUserEndpoint(svc)),
		DeleteUserEndpoint: middleware.LoggingMiddleware(logger)(endpoint.MakeDeleteUserEndpoint(svc)),
	}

	handler := transport.NewHTTPHandler(endpoints)

	log.Fatal(http.ListenAndServe(":8080", handler))

}
