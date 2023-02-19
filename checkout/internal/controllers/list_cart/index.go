package listcart

import (
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic *domain.Model
}

type Request struct {
	User int64 `json:"user"`
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

type ResponseItem struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Response struct {
	Items []ResponseItem `json:"items"`
	Total uint64         `json:"total"`
}

func New(domainLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: domainLogic,
	}
}

func (h *Handler) Handle(req Request) (Response, error) {
	res, err := h.businessLogic.ListCart(uint32(req.User))

	if err != nil {
		return Response{}, errors.Wrap(err, "getting failed")
	}

	items := make([]ResponseItem, len(res.Offers))

	for i, offer := range res.Offers {
		items[i] = ResponseItem{
			Sku:   offer.Sku,
			Count: offer.Count,
			Name:  offer.Name,
			Price: uint32(offer.Price),
		}
	}

	return Response{
		Items: items,
		Total: uint64(res.TotalPrice),
	}, nil
}
