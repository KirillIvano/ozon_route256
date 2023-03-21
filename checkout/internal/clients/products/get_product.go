package products_client

import (
	"context"
	"fmt"
	"route256/checkout/internal/domain"
	productsService "route256/products/pkg/products_service"
	"time"
)

type GetProductRequest struct {
	Sku   uint32 `json:"sku"`
	Token string `json:"token"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *client) GetProduct(ctx context.Context, sku uint32) (domain.ProductInfo, error) {
	start := time.Now()
	now := c.rateLimiter.Take()
	fmt.Printf("Request for sku=%d blocked by limiter for %d ms\n", sku, now.Sub(start)/time.Millisecond)

	res, err := c.client.GetProduct(ctx, &productsService.GetProductRequest{Sku: sku, Token: c.token})

	if err != nil {
		return domain.ProductInfo{}, err
	}

	return domain.ProductInfo{
		Price: res.GetPrice(),
		Name:  res.GetName(),
	}, nil
}
