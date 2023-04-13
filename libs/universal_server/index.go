package universal_server

import (
	"context"
	"fmt"
	"net"
	"route256/libs/logger"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UniversalServer struct {
	server         *grpc.Server
	port           int32
	metricAddr     string
	metricsManager MetricsManager
}

func (s *UniversalServer) GetServerRegistrar() grpc.ServiceRegistrar {
	return s.server
}

func (s *UniversalServer) Listen(ctx context.Context) {
	netListener := net.ListenConfig{}
	listener, err := netListener.Listen(ctx, "tcp", fmt.Sprintf(":%d", s.port))

	if err != nil {
		logger.Fatal("failed to listen: ", zap.Error(err))
	}

	serverDone := make(chan struct{})
	defer close(serverDone)

	s.metricsManager.startMetricsServer(ctx, s.metricAddr)

	go func() {
		defer func() { serverDone <- struct{}{} }()

		if err := s.server.Serve(listener); err != nil {
			logger.Fatal("failed to serve: ", zap.Error(err))
		}
	}()

	for {
		select {
		case <-serverDone:
			return
		case <-ctx.Done():
			s.server.GracefulStop()
		}
	}
}

func New(project string, port int32, metricAddr string) *UniversalServer {
	loggerInterceptor := createLoggerInterceptor()
	metricsManager := newMetricsManager(project)

	handler := grpc.ChainUnaryInterceptor(
		loggerInterceptor,
		metricsManager.metricsInterceptor,
	)

	grpcServer := grpc.NewServer(handler, grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
	))

	reflection.Register(grpcServer)

	return &UniversalServer{
		server:         grpcServer,
		port:           port,
		metricAddr:     metricAddr,
		metricsManager: *metricsManager,
	}
}
