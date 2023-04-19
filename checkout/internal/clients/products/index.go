package products_client

import (
	"context"
	"route256/checkout/internal/domain"
	"route256/checkout/pkg/rate_limiter"
	"route256/libs/cache"
	"route256/libs/logger"
	productsService "route256/products/pkg/products_service"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	token string

	cache       *cache.Cache[domain.ProductInfo]
	rateLimiter rate_limiter.RateLimiter
	client      productsService.ProductServiceClient
	conn        *grpc.ClientConn
}

type ProductsClient interface {
	GetProduct(ctx context.Context, sku uint32) (domain.ProductInfo, error)
}

func (c *client) Close() {
	c.conn.Close()
}

const RpsLimit = 10

func New(ctx context.Context, urlOrigin string, token string) *client {
	conn, err := grpc.DialContext(
		ctx,
		urlOrigin,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)

	if err != nil {
		logger.Fatal("failed to connect to server:", zap.Error(err))
	}

	c := productsService.NewProductServiceClient(conn)
	rateLimiter := rate_limiter.New(RpsLimit)
	cache := cache.NewCache[domain.ProductInfo](time.Second * 30)

	return &client{
		cache:       cache,
		token:       token,
		client:      c,
		conn:        conn,
		rateLimiter: *rateLimiter,
	}
}
