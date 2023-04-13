package main

import (
	"context"
	"flag"
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
	"route256/libs/universal_server"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	metricAddr = flag.String("addr", ":9080", "the port to listen")
	develMode  = flag.Bool("devel", false, "development mode")
)

func main() {
	ctx := context.Background()

	ctx, stopSignalListen := signal.NotifyContext(ctx, os.Interrupt)
	defer stopSignalListen()

	logger.Init(*develMode)
	tracing.Init("checkout")

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

	server := universal_server.New("checkout", config.ConfigData.Port, *metricAddr)

	checkoutService.RegisterCheckoutServer(server.GetServerRegistrar(), checkoutServer.New(businessLogic))

	server.Listen(ctx)
}
