package products_client

import (
	"context"
	"log"
	"route256/checkout/pkg/rate_limiter"
	productsService "route256/products/pkg/products_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	token string

	rateLimiter rate_limiter.RateLimiter
	client      productsService.ProductServiceClient
	conn        *grpc.ClientConn
}

func (c *client) Close() {
	c.conn.Close()
}

const RpsLimit = 10

func New(ctx context.Context, urlOrigin string, token string) *client {
	conn, err := grpc.DialContext(ctx, urlOrigin, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	c := productsService.NewProductServiceClient(conn)
	rateLimiter := rate_limiter.New(RpsLimit)

	return &client{
		token:       token,
		client:      c,
		conn:        conn,
		rateLimiter: *rateLimiter,
	}
}
