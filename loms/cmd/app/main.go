package main

import (
	"fmt"
	"log"
	"net"
	"route256/loms/internal/domain"
	lomsServer "route256/loms/internal/loms_server_v1"
	lomsV1 "route256/loms/pkg/loms_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = "8081"

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	domain := domain.New()

	grpcServer := grpc.NewServer()
	lomsV1.RegisterLomsV1Server(grpcServer, lomsServer.New(domain))
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
