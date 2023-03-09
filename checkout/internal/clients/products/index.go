package products_client

import (
	"context"
	"log"
	productsService "route256/products/pkg/products_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	token string

	client productsService.ProductServiceClient
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

	c := productsService.NewProductServiceClient(conn)

	return &client{
		token:  token,
		client: c,
		conn:   conn,
	}
}
