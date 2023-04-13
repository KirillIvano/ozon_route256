package universal_server

import (
	"context"
	"route256/libs/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func createLoggerInterceptor() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger.Info("incoming grpc request", zap.String("method", info.FullMethod))

		return handler(ctx, req)
	}
}
