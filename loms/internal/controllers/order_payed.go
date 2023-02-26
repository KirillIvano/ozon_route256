package controllers

import (
	"github.com/pkg/errors"
)

type OrderPayedRequest struct {
	OrderId int64 `json:"orderId"`
}
type OrderPayedResponse struct{}

var (
	ErrOrderPayedEmptyOrderId = errors.New("empty order id")
)

func (r OrderPayedRequest) Validate() error {
	if r.OrderId == 0 {
		return ErrOrderPayedEmptyOrderId
	}

	return nil
}

func (h *LomsHandlersRegistry) HandleOrderPayed(req OrderPayedRequest) (OrderPayedResponse, error) {
	err := h.domainLogic.SetOrderPayed(req.OrderId)

	if err != nil {
		return OrderPayedResponse{}, errors.Wrap(err, "cancellation failed")
	}

	return OrderPayedResponse{}, nil
}
