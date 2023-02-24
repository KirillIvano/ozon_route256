package createorder

import (
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic *domain.Model
}

type Request struct {
	User  int64 `json:"user"`
	Items []struct {
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	} `json:"items"`
}

var (
	ErrEmptyUser = errors.New("empty user")
)

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}

	return nil
}

type Response struct {
	OrderId int64 `json:"orderId"`
}

func New(domainLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: domainLogic,
	}
}

func (h *Handler) Handle(req Request) (Response, error) {
	items := make([]domain.OrderItem, len(req.Items))

	for idx, item := range req.Items {
		items[idx] = domain.OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	orderId, err := h.businessLogic.CreateOrder(req.User, items)

	if err != nil {
		return Response{}, errors.Wrap(err, "creation failed")
	}

	return Response{OrderId: orderId}, nil
}
