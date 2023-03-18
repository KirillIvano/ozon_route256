package main

import (
	"context"
	"fmt"
	"log"
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

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startServer(ctx context.Context, businessLogic *domain.CheckoutDomain) *grpc.Server {
	netListener := net.ListenConfig{}
	listener, err := netListener.Listen(ctx, "tcp", fmt.Sprintf(":%d", config.ConfigData.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()

	checkoutService.RegisterCheckoutServer(grpcServer, checkoutServer.New(businessLogic))
	reflection.Register(grpcServer)

	serverDone := make(chan struct{})
	defer close(serverDone)

	go func() {
		defer func() { serverDone <- struct{}{} }()

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("failed to serve: ", err)
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

	wp := workerPool.New(ctx, 5)
	defer wp.GracefulClose()

	err := config.Init()
	if err != nil {
		log.Fatal("config init failed")
	}

	dbConfig, err := pgxpool.ParseConfig(config.ConfigData.Services.Database)
	if err != nil {
		log.Fatal("database config parse failed")
	}

	conn, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		log.Fatal(err)
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
