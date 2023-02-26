package controllers

import (
	"github.com/pkg/errors"
)

type DeleteFromCartRequest struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

var (
	ErrDeleteFromCartEmptyUser  = errors.New("empty user")
	ErrDeleteFromCartEmptySKU   = errors.New("empty sku")
	ErrDeleteFromCartBadRequest = errors.New("bad request")
)

func (r DeleteFromCartRequest) Validate() error {
	if r.User == 0 {
		return ErrDeleteFromCartEmptyUser
	}

	if r.Sku == 0 {
		return ErrDeleteFromCartEmptySKU
	}

	if r.Count <= 0 {
		return ErrDeleteFromCartBadRequest
	}

	return nil
}

type DeleteFromCartResponse struct{}

func (h *CheckoutHandlersRegistry) HandleDeleteFromCart(req DeleteFromCartRequest) (DeleteFromCartResponse, error) {
	err := h.domainLogic.DeleteFromCart(req.User, req.Sku, req.Count)

	if err != nil {
		return DeleteFromCartResponse{}, errors.Wrap(err, "deletion failed")
	}

	return DeleteFromCartResponse{}, nil
}
