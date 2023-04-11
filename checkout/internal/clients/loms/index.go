package loms_client

import (
	"context"
	"route256/checkout/internal/domain"
	"route256/libs/logger"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
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

type LomsClient interface {
	CreateOrder(ctx context.Context, user int64, items []domain.CartItem) (int64, error)
	Stocks(ctx context.Context, sku uint32) ([]domain.Stock, error)
}

func New(ctx context.Context, urlOrigin string) *Client {
	conn, err := grpc.DialContext(
		ctx,
		urlOrigin,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)

	if err != nil {
		logger.Fatal("failed to connect to server: %v", zap.Error(err))
	}

	c := lomsService.NewLomsClient(conn)

	return &Client{
		client: c,
		conn:   conn,
	}
}
