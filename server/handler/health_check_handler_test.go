package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	echov1 "github.com/ohmygrpc/idl/gen/go/services/echo/v1"
)

func TestHealthCheck(t *testing.T) {
	ctx := context.Background()
	req := &echov1.HealthCheckRequest{}

	resp, err := HealthCheck()(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, &echov1.HealthCheckResponse{}, resp)
}
