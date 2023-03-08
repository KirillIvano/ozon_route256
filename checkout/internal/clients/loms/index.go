package loms_client

import (
	"context"
	"log"
	lomsService "route256/loms/pkg/loms_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client lomsService.LomsClient
	conn   *grpc.ClientConn
}

func (c *Client) Close() {
	c.conn.Close()
}

func New(ctx context.Context, urlOrigin string) *Client {
	conn, err := grpc.DialContext(ctx, urlOrigin, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	c := lomsService.NewLomsClient(conn)

	return &Client{
		client: c,
		conn:   conn,
	}
}
