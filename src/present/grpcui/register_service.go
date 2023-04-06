package grpcui

import (
	"base/src/present/grpcui/middlewares"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type GrpcServerIn struct {
	fx.In
	Server *grpc.Server
}

func NewGRPCServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middlewares.TrackingInterceptor(),
			middlewares.LogInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(), // recover from gRPC handler panics into a gRPC error with `code.Internal`
		)),
	)
	return server
}

func registerGRPCServices(in GrpcServerIn) {

}
