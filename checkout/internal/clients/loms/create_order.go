package loms_client

import (
	"context"
	"route256/checkout/internal/domain"
	lomsService "route256/loms/pkg/loms_service"
)

func (c *Client) CreateOrder(ctx context.Context, user int64, items []domain.CartItem) (int64, error) {
	createOrderItems := make([]*lomsService.OrderItem, len(items))
	for idx, item := range items {
		createOrderItems[idx] = &lomsService.OrderItem{
			Sku:   item.Sku,
			Count: uint32(item.Count),
		}
	}

	requestData := lomsService.CreateOrderParams{
		User:  user,
		Items: createOrderItems,
	}

	response, err := c.client.CreateOrder(ctx, &requestData)

	if err != nil {
		return 0, err
	}

	return response.OrderId, nil
}
