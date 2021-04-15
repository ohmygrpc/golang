package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/ohmygrpc/golang/config"
	echov1 "github.com/ohmygrpc/idl/gen/go/services/echo/v1"
)

func NewHTTPServer(ctx context.Context, cfg config.Config) (*http.Server, error) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		),
	)
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := echov1.RegisterEchoServiceHandlerFromEndpoint(
		ctx,
		mux,
		cfg.Setting().GRPCServerEndpoint,
		options,
	); err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    ":" + cfg.Setting().HTTPServerPort,
		Handler: mux,
	}

	return server, nil
}
