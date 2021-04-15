package handler

import (
	"context"

	echov1 "github.com/ohmygrpc/idl/gen/go/services/echo/v1"
)

type HealthCheckHandlerFunc func(ctx context.Context, req *echov1.HealthCheckRequest) (*echov1.HealthCheckResponse, error)

func HealthCheck() HealthCheckHandlerFunc {
	return func(ctx context.Context, req *echov1.HealthCheckRequest) (*echov1.HealthCheckResponse, error) {
		return &echov1.HealthCheckResponse{}, nil
	}
}
