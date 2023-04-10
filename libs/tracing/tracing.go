package tracing

import (
	"context"
	"route256/libs/logger"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Init(serviceName string) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("Cannot init tracing", zap.Error(err))
	}
}

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, info.FullMethod)
	defer span.Finish()

	span.SetTag("method", info.FullMethod)

	if spanContext, ok := span.Context().(jaeger.SpanContext); ok {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", spanContext.TraceID().String())
	}

	res, err := handler(ctx, req)

	if err != nil {
		ext.Error.Set(span, true)
	}

	code, isStatus := status.FromError(err)

	if !isStatus {
		span.SetTag("status_code", "unknown")
	} else {
		stringStatus := strconv.Itoa(int(code.Code()))
		span.SetTag("status_code", stringStatus)
	}

	return res, err
}
