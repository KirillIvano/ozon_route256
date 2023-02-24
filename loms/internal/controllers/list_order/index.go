package listorder

import (
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic *domain.Model
}

type Request struct {
	OrderId int64 `json:"orderId"`
}

var (
	ErrEmptyOrderId = errors.New("empty order id")
)

func (r Request) Validate() error {
	if r.OrderId == 0 {
		return ErrEmptyOrderId
	}

	return nil
}

type ResponseOrderItem = struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	Status string              `json:"status"`
	User   int64               `json:"user"`
	Items  []ResponseOrderItem `json:"items"`
}

func New(domainLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: domainLogic,
	}
}

func (h *Handler) Handle(req Request) (Response, error) {
	orderInfo, err := h.businessLogic.ListOrder(req.OrderId)

	if err != nil {
		return Response{}, errors.Wrap(err, "creation failed")
	}

	responseItems := make([]ResponseOrderItem, len(orderInfo.Items))
	for idx, item := range orderInfo.Items {
		responseItems[idx] = ResponseOrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	return Response{
		Status: orderInfo.Status,
		User:   orderInfo.User,
		Items:  responseItems,
	}, nil
}
