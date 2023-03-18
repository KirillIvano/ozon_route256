package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"route256/loms/internal/config"
	"route256/loms/internal/domain"
	lomsServer "route256/loms/internal/loms_server"
	"route256/loms/internal/repository"
	statusActualizer "route256/loms/internal/status_actualizer"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startListen(ctx context.Context, domain *domain.LomsDomain) {
	netListener := net.ListenConfig{}
	listener, err := netListener.Listen(ctx, "tcp", fmt.Sprintf(":%d", config.ConfigData.Port))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	lomsService.RegisterLomsServer(grpcServer, lomsServer.New(domain))
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}

func main() {
	ctx := context.Background()
	ctx, stopSignalListen := signal.NotifyContext(ctx, os.Interrupt)
	defer stopSignalListen()

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		log.Fatal("config init failed")
	}

	conn, err := pgxpool.New(ctx, config.ConfigData.Services.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("database connected successfully")

	repository := repository.Connect(conn)
	domain := domain.New(repository)

	enf := statusActualizer.New(ctx, domain)

	enf.Start()
	defer enf.Close()

	startListen(ctx, domain)
}
