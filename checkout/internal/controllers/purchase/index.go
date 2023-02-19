package purchase

import (
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic *domain.Model
}

type Request struct {
	User uint64 `json:"user"`
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

type Response struct{}

func New(domainLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: domainLogic,
	}
}

func (h *Handler) Handle(req Request) (Response, error) {
	err := h.businessLogic.Purchase(req.User)

	if err != nil {
		return Response{}, errors.Wrap(err, "purchase failed")
	}

	return Response{}, nil
}
