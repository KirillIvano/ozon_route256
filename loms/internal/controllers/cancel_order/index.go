package cancelorder

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

type Response struct{}

func New(domainLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: domainLogic,
	}
}

func (h *Handler) Handle(req Request) (Response, error) {
	err := h.businessLogic.CancelOrder(req.OrderId)

	if err != nil {
		return Response{}, errors.Wrap(err, "cancellation failed")
	}

	return Response{}, nil
}
