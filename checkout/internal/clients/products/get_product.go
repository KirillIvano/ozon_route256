package products_client

import (
	"context"
	"route256/checkout/internal/domain"
	productsService "route256/products/pkg/products_service"
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
	res, err := c.client.GetProduct(ctx, &productsService.GetProductRequest{Sku: sku, Token: c.token})

	if err != nil {
		return domain.ProductInfo{}, err
	}

	return domain.ProductInfo{
		Price: res.GetPrice(),
		Name:  res.GetName(),
	}, nil
}
