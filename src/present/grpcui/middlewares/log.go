package middlewares

import (
	"base/src/common/log"
	"context"
	"google.golang.org/grpc"
)

func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		log.Info(ctx, "Receive request")
		return resp, err
	}
}
