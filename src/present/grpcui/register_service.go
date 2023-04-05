package grpcui

import (
	"base/src/common/configs"
	"base/src/present/grpcui/middlewares"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GrpcServerIn struct {
	fx.In
	Server *grpc.Server
}

func BuildGrpcModules() fx.Option {
	return fx.Options(
		fx.Provide(newGRPCServer),

		fx.Invoke(registerGRPCServices),
	)
}

func newGRPCServer(log *zap.SugaredLogger, lc fx.Lifecycle) (*grpc.Server, error) {
	server := grpc.NewServer(RegisterOptsServer()...)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Debug("OnStart GRPC Server")
			lis, err := net.Listen("tcp", configs.Get().Server.Grpc.Address)
			if err != nil {
				return err
			}
			go func() {
				log.Infof("GRPC Server is listening at %v", lis.Addr())
				if err = server.Serve(lis); err != nil {
					log.Fatal("Listen GRPC Server error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Stop()
			return nil
		},
	})

	return server, nil
}

func RegisterOptsServer() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middlewares.TrackingInterceptor(),
			middlewares.LogInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(), // recover from gRPC handler panics into a gRPC error with `code.Internal`
		)),
	}
}

func registerGRPCServices(in GrpcServerIn) {

}
