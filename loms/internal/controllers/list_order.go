package controllers

import (
	"github.com/pkg/errors"
)

type ListOrderRequest struct {
	OrderId int64 `json:"orderId"`
}

var (
	ErrListOrderEmptyOrderId = errors.New("empty order id")
)

func (r ListOrderRequest) Validate() error {
	if r.OrderId == 0 {
		return ErrListOrderEmptyOrderId
	}

	return nil
}

type ListOrderResponseItem = struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type ListOrderResponse struct {
	Status string                  `json:"status"`
	User   int64                   `json:"user"`
	Items  []ListOrderResponseItem `json:"items"`
}

func (h *LomsHandlersRegistry) HandleListOrder(req ListOrderRequest) (ListOrderResponse, error) {
	orderInfo, err := h.domainLogic.ListOrder(req.OrderId)

	if err != nil {
		return ListOrderResponse{}, errors.Wrap(err, "creation failed")
	}

	responseItems := make([]ListOrderResponseItem, len(orderInfo.Items))
	for idx, item := range orderInfo.Items {
		responseItems[idx] = ListOrderResponseItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	return ListOrderResponse{
		Status: orderInfo.Status,
		User:   orderInfo.User,
		Items:  responseItems,
	}, nil
}
