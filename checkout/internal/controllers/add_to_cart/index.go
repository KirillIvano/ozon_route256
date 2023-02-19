package addtocart

import (
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic *domain.Model
}

type Request struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

var (
	ErrEmptyUser  = errors.New("empty user")
	ErrEmptySKU   = errors.New("empty sku")
	ErrBadRequest = errors.New("bad request")
)

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}

	if r.Sku == 0 {
		return ErrEmptySKU
	}

	if r.Count <= 0 {
		return ErrBadRequest
	}

	return nil
}

type Response struct {
	Result string
}

func New(domainLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: domainLogic,
	}
}

func (h *Handler) Handle(req Request) (Response, error) {
	err := h.businessLogic.AddToCart(req.User, req.Sku, req.Count)

	if err != nil {
		return Response{}, errors.Wrap(err, "creation failed")
	}

	return Response{Result: "succeed"}, nil
}
