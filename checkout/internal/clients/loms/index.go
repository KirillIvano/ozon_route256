package loms_client

import (
	"context"
	"log"
	lomsV1 "route256/loms/pkg/loms_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client lomsV1.LomsV1Client
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

	c := lomsV1.NewLomsV1Client(conn)

	return &Client{
		client: c,
		conn:   conn,
	}
}
