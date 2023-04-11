package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	checkoutServer "route256/checkout/internal/checkout_server"
	lomsClient "route256/checkout/internal/clients/loms"
	productsClient "route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/repository"
	checkoutService "route256/checkout/pkg/checkout_service"
	workerPool "route256/checkout/pkg/worker_pool"
	"route256/libs/logger"
	"route256/libs/tracing"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	metricAddr = flag.String("addr", ":9080", "the port to listen")
	develMode  = flag.Bool("devel", false, "development mode")
)

func startServer(ctx context.Context, businessLogic *domain.CheckoutDomain) *grpc.Server {
	netListener := net.ListenConfig{}
	listener, err := netListener.Listen(ctx, "tcp", fmt.Sprintf(":%d", config.ConfigData.Port))
	if err != nil {
		logger.Fatal("failed to listen server", zap.Error(err))
	}

	handler := grpc.ChainUnaryInterceptor(
		logger.Interceptor,
		otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
	)
	grpcServer := grpc.NewServer(handler)

	checkoutService.RegisterCheckoutServer(grpcServer, checkoutServer.New(businessLogic))
	reflection.Register(grpcServer)

	serverDone := make(chan struct{})
	defer close(serverDone)

	go func() {
		defer func() { serverDone <- struct{}{} }()

		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	// ждем закрытия окончания работы сервера
	for {
		select {
		case <-serverDone:
			return grpcServer
		case <-ctx.Done():
			grpcServer.GracefulStop()
		}
	}
}

func main() {
	ctx := context.Background()

	ctx, stopSignalListen := signal.NotifyContext(ctx, os.Interrupt)
	defer stopSignalListen()

	logger.Init(*develMode)
	tracing.Init("checkout_service")

	wp := workerPool.New(ctx, 5)
	defer wp.GracefulClose()

	err := config.Init()
	if err != nil {
		logger.Fatal("config init failed")
	}

	dbConfig, err := pgxpool.ParseConfig(config.ConfigData.Services.Database)
	if err != nil {
		logger.Fatal("database config parse failed")
	}

	conn, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		logger.Fatal("pool init failed")
	}
	defer conn.Close()

	repository := repository.New(conn)

	lomsClient := lomsClient.New(ctx, config.ConfigData.Services.Loms)
	defer lomsClient.Close()

	productClient := productsClient.New(ctx, config.ConfigData.Services.Products, config.ConfigData.Token)
	defer productClient.Close()

	businessLogic := domain.New(lomsClient, productClient, repository, wp)

	startServer(ctx, businessLogic)
}
