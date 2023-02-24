package loms_client

import (
	"net/http"
	"route256/checkout/internal/domain"
	"route256/libs/jsonreqwrap"
)

type CreateOrderItem struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderRequestBody struct {
	User  uint64            `json:"user"`
	Items []CreateOrderItem `json:"items"`
}

type CreateOrderResponse struct {
	OrderId uint64 `json:"orderId"`
}

func (c *Client) CreateOrder(user uint64, items []domain.CartItem) (uint64, error) {
	reqClient := jsonreqwrap.NewClient[CreateOrderRequestBody, CreateOrderResponse](
		c.urlCreateOrder,
		http.MethodPost,
	)

	createOrderItems := make([]CreateOrderItem, len(items))
	for idx, item := range items {
		createOrderItems[idx] = CreateOrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	requestData := CreateOrderRequestBody{
		User:  user,
		Items: createOrderItems,
	}

	response, err := reqClient.Run(requestData)

	if err != nil {
		return 0, err
	}

	return response.OrderId, nil
}
