package bootstrap

import (
	"base/src/common/configs"
	"base/src/common/log"
	"base/src/present/grpcui"
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func BuildGrpcModules() fx.Option {
	return fx.Options(
		fx.Provide(grpcui.NewGRPCServer),

		fx.Invoke(newGRPCServer),
	)
}

func newGRPCServer(lc fx.Lifecycle, server *grpc.Server) {
	logger := log.GetLogger().GetZap()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Debug("OnStart GRPC Server")
			lis, err := net.Listen("tcp", configs.Get().Server.Grpc.Address)
			if err != nil {
				return err
			}
			go func() {
				logger.Infof("GRPC Server is listening at %v", lis.Addr())
				if err = server.Serve(lis); err != nil {
					logger.Fatal("Listen GRPC Server error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Stop()
			return nil
		},
	})
}
