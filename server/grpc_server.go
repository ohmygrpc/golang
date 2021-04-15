package server

import (
	"context"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/ohmygrpc/golang/config"
	"github.com/ohmygrpc/golang/server/handler"
	echov1 "github.com/ohmygrpc/idl/gen/go/services/echo/v1"
)

type echoServiceServer struct {
	echov1.EchoServiceServer

	cfg config.Config
}

func NewEchoServiceServer(cfg config.Config) (*echoServiceServer, error) {
	return &echoServiceServer{cfg: cfg}, nil
}

func (s *echoServiceServer) Config() config.Config {
	return s.cfg
}

func (s *echoServiceServer) RegisterServer(srv *grpc.Server) {
	echov1.RegisterEchoServiceServer(srv, s)
}

func (s *echoServiceServer) HealthCheck(ctx context.Context, req *echov1.HealthCheckRequest) (*echov1.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}

func (s *echoServiceServer) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	return handler.Echo()(ctx, req)
}

func NewGRPCServer(cfg config.Config) (*grpc.Server, error) {
	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 30 * time.Second,
		}),
	)

	echoServer, err := NewEchoServiceServer(cfg)
	if err != nil {
		return nil, err
	}

	echov1.RegisterEchoServiceServer(grpcServer, echoServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}
