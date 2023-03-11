package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"route256/loms/internal/domain"
	lomsServer "route256/loms/internal/loms_server"
	"route256/loms/internal/repository"
	lomsService "route256/loms/pkg/loms_service"

	pgx "github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = "8081"
const CONNECTION_STRING = "postgres://user:password@localhost:8091/loms?sslmode=disable"

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	log.Println("database connected successfully")

	repository := repository.Connect(conn)
	domain := domain.New(repository)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
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
