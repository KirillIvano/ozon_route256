package main

import (
	"context"
	"fmt"
	"log"
	"net"
	checkoutServer "route256/checkout/internal/checkout_server"
	lomsClient "route256/checkout/internal/clients/loms"
	productsClient "route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/repository"
	checkoutService "route256/checkout/pkg/checkout_service"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = "8080"

func main() {
	ctx := context.Background()

	err := config.Init()
	if err != nil {
		log.Fatal("config init failed")
	}

	conn, err := pgx.Connect(ctx, config.ConfigData.Services.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	log.Println("database connected successfully")

	repository := repository.New(conn)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lomsClient := lomsClient.New(ctx, config.ConfigData.Services.Loms)
	defer lomsClient.Close()

	productClient := productsClient.New(ctx, config.ConfigData.Services.Products, config.ConfigData.Token)
	defer productClient.Close()

	businessLogic := domain.New(lomsClient, productClient, repository)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	checkoutService.RegisterCheckoutServer(grpcServer, checkoutServer.New(businessLogic))
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
