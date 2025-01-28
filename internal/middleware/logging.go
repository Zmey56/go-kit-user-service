package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			start := time.Now()
			response, err := next(ctx, request)
			logger.Log(
				"timestamp", time.Now().Format(time.RFC3339),
				"took", time.Since(start),
				"err", err,
			)
			return response, err
		}
	}
}
