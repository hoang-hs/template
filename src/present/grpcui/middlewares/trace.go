package middlewares

import (
	"base/src/core/constant"
	"context"
	"encoding/hex"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"math/rand"
)

func TrackingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		spanContext := trace.SpanContextFromContext(ctx)
		traceId := ""
		if spanContext.SpanID().IsValid() {
			traceId = spanContext.TraceID().String()
		} else {
			traceIdByte := make([]byte, 16)
			rand.Read(traceIdByte)
			traceId = hex.EncodeToString(traceIdByte[:])
		}
		newCtx := context.WithValue(ctx, constant.TraceIdName, traceId)
		resp, err = handler(newCtx, req)
		return resp, err
	}
}
