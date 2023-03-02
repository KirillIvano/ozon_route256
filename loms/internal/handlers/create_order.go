package handlers

import (
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
)

type CreateOrderRequest struct {
	User  int64 `json:"user"`
	Items []struct {
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	} `json:"items"`
}

var (
	ErrEmptyUser = errors.New("empty user")
)

func (r CreateOrderRequest) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}

	return nil
}

type CreateOrderResponse struct {
	OrderId int64 `json:"orderId"`
}

func (h *LomsHandlersRegistry) CreateOrder(req CreateOrderRequest) (CreateOrderResponse, error) {
	items := make([]domain.OrderItem, len(req.Items))

	for idx, item := range req.Items {
		items[idx] = domain.OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	orderId, err := h.domainLogic.CreateOrder(req.User, items)

	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "creation failed")
	}

	return CreateOrderResponse{OrderId: orderId}, nil
}
