package handlers

import (
	"github.com/pkg/errors"
)

type CancelOrderRequest struct {
	OrderId int64 `json:"orderId"`
}

var (
	ErrCancelOrderEmptyOrderId = errors.New("empty order id")
)

func (r CancelOrderRequest) Validate() error {
	if r.OrderId == 0 {
		return ErrCancelOrderEmptyOrderId
	}

	return nil
}

type CancelOrderResponse struct{}

func (h *LomsHandlersRegistry) CancelOrder(req CancelOrderRequest) (CancelOrderResponse, error) {
	err := h.domainLogic.CancelOrder(req.OrderId)

	if err != nil {
		return CancelOrderResponse{}, errors.Wrap(err, "cancellation failed")
	}

	return CancelOrderResponse{}, nil
}
