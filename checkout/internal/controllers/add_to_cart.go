package controllers

import (
	"github.com/pkg/errors"
)

type AddToCartRequest struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

var (
	ErrAddToCartEmptyUser  = errors.New("empty user")
	ErrAddToCartEmptySKU   = errors.New("empty sku")
	ErrAddToCartBadRequest = errors.New("bad request")
)

func (r AddToCartRequest) Validate() error {
	if r.User == 0 {
		return ErrAddToCartEmptyUser
	}

	if r.Sku == 0 {
		return ErrAddToCartEmptySKU
	}

	if r.Count <= 0 {
		return ErrAddToCartBadRequest
	}

	return nil
}

type AddToCartResponse struct {
	Result string
}

func (h *CheckoutHandlersRegistry) HandleAddToCart(req AddToCartRequest) (AddToCartResponse, error) {
	err := h.domainLogic.AddToCart(req.User, req.Sku, req.Count)

	if err != nil {
		return AddToCartResponse{}, errors.Wrap(err, "creation failed")
	}

	return AddToCartResponse{Result: "succeed"}, nil
}
