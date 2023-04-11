package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"route256/libs/kafka"
	"route256/libs/logger"
	"route256/libs/tracing"
	"route256/loms/internal/config"
	"route256/loms/internal/domain"
	lomsServer "route256/loms/internal/loms_server"
	"route256/loms/internal/metrics"
	orderSender "route256/loms/internal/order_sender"
	"route256/loms/internal/repository"
	statusActualizer "route256/loms/internal/status_actualizer"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	metricAddr = flag.String("addr", ":9081", "the port to listen")
	develMode  = flag.Bool("devel", false, "development mode")
)

func startServer(ctx context.Context, businessLogic *domain.LomsDomain) {
	netListener := net.ListenConfig{}
	listener, err := netListener.Listen(ctx, "tcp", fmt.Sprintf(":%d", config.ConfigData.Port))
	if err != nil {
		logger.Fatal("failed to listen: ", zap.Error(err))
	}

	handler := grpc.ChainUnaryInterceptor(
		logger.Interceptor,
		metrics.Interceptor,
	)
	grpcServer := grpc.NewServer(handler, grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
	),
	)

	lomsService.RegisterLomsServer(grpcServer, lomsServer.New(businessLogic))
	reflection.Register(grpcServer)

	serverDone := make(chan struct{})
	defer close(serverDone)

	go func() {
		defer func() { serverDone <- struct{}{} }()

		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatal("failed to serve: ", zap.Error(err))
		}
	}()

	// ждем закрытия окончания работы сервера
	for {
		select {
		case <-serverDone:
			return
		case <-ctx.Done():
			grpcServer.GracefulStop()
		}
	}
}

func initMetrics(ctx context.Context) {
	srv := http.Server{Addr: *metricAddr}

	metric := metrics.New()
	metric.Observe(ctx)

	http.Handle("/metrics", metric.Handler())

	go func() {
		err := srv.ListenAndServe()

		if err != nil {
			logger.Fatal("metric server failed to start", zap.Error(err))
		}
	}()
}

func main() {
	ctx := context.Background()
	ctx, stopSignalListen := signal.NotifyContext(ctx, os.Interrupt)
	defer stopSignalListen()

	logger.Init(*develMode)
	tracing.Init("loms_service")

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		logger.Fatal("config init failed", zap.Error(err))
	}

	initMetrics(ctx)

	conn, err := pgxpool.New(ctx, config.ConfigData.Services.Database)
	if err != nil {
		logger.Fatal("pool connection error", zap.Error(err))
	}
	defer conn.Close()

	logger.Info("database connected successfully")

	producer, err := kafka.NewSyncProducer(config.ConfigData.Brokers)
	if err != nil {
		logger.Fatal("kafka connection failed", zap.Error(err))
	}
	orderSender := orderSender.NewOrderSender(producer, config.ConfigData.OrderTopic)

	repository := repository.Connect(conn)
	domain := domain.New(repository, orderSender)

	enf := statusActualizer.New(ctx, domain)

	enf.Start()
	defer enf.Close()

	startServer(ctx, domain)
}
