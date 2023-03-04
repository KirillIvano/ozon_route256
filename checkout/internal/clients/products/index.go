package products_client

import (
	"context"
	"log"
	"route256/products/pkg/products_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	token string

	client products_v1.ProductServiceClient
	conn   *grpc.ClientConn
}

func (c *client) Close() {
	c.conn.Close()
}

func New(ctx context.Context, urlOrigin string, token string) *client {
	conn, err := grpc.DialContext(ctx, urlOrigin, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	c := products_v1.NewProductServiceClient(conn)

	return &client{
		token:  token,
		client: c,
		conn:   conn,
	}
}
