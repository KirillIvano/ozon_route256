package handlers

import (
	"github.com/pkg/errors"
)

type PurchaseRequest struct {
	User uint64 `json:"user"`
}

var (
	ErrPurchaseEmptyUser = errors.New("empty user")
)

func (r PurchaseRequest) Validate() error {
	if r.User == 0 {
		return ErrPurchaseEmptyUser
	}

	return nil
}

type PurchaseResponse struct{}

func (h *CheckoutHandlersRegistry) Purchase(req PurchaseRequest) (PurchaseResponse, error) {
	err := h.domainLogic.Purchase(req.User)

	if err != nil {
		return PurchaseResponse{}, errors.Wrap(err, "purchase failed")
	}

	return PurchaseResponse{}, nil
}
