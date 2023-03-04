package loms_client

import (
	"context"
	"route256/checkout/internal/domain"
	lomsV1 "route256/loms/pkg/loms_v1"
)

func (c *Client) CreateOrder(ctx context.Context, user int64, items []domain.CartItem) (int64, error) {
	createOrderItems := make([]*lomsV1.OrderItem, len(items))
	for idx, item := range items {
		createOrderItems[idx] = &lomsV1.OrderItem{
			Sku:   item.Sku,
			Count: uint32(item.Count),
		}
	}

	requestData := lomsV1.CreateOrderParams{
		User:  user,
		Items: createOrderItems,
	}

	response, err := c.client.CreateOrder(ctx, &requestData)

	if err != nil {
		return 0, err
	}

	return response.OrderId, nil
}
