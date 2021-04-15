package handler

import (
	"context"

	echov1 "github.com/ohmygrpc/idl/gen/go/services/echo/v1"
)

type EchoHandlerFunc func(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error)

func Echo() EchoHandlerFunc {
	return func(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
		return &echov1.EchoResponse{
			Msg: req.Msg,
		}, nil
	}
}
