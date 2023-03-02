package handlers

import (
	"github.com/pkg/errors"
)

type ListCartRequest struct {
	User int64 `json:"user"`
}

var (
	ErrListCartEmptyUser = errors.New("empty user")
)

func (r ListCartRequest) Validate() error {
	if r.User == 0 {
		return ErrListCartEmptyUser
	}

	return nil
}

type ListCartResponseItem struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type ListCartResponse struct {
	Items []ListCartResponseItem `json:"items"`
	Total uint64                 `json:"total"`
}

func (h *CheckoutHandlersRegistry) ListCart(req ListCartRequest) (ListCartResponse, error) {
	res, err := h.domainLogic.ListCart(uint32(req.User))

	if err != nil {
		return ListCartResponse{}, errors.Wrap(err, "getting failed")
	}

	items := make([]ListCartResponseItem, len(res.Offers))

	for i, offer := range res.Offers {
		items[i] = ListCartResponseItem{
			Sku:   offer.Sku,
			Count: offer.Count,
			Name:  offer.Name,
			Price: uint32(offer.Price),
		}
	}

	return ListCartResponse{
		Items: items,
		Total: uint64(res.TotalPrice),
	}, nil
}
