package stocks

import (
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic *domain.Model
}

type Request struct {
	Sku uint32 `json:"sku"`
}

var (
	ErrEmptySku = errors.New("empty sku")
)

func (r Request) Validate() error {
	if r.Sku == 0 {
		return ErrEmptySku
	}

	return nil
}

type ResponseStock = struct {
	WarehouseID int64  `json:"warehouseId"`
	Count       uint64 `json:"count"`
}

type Response struct {
	Stocks []ResponseStock `json:"stocks"`
}

func New(domainLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: domainLogic,
	}
}

func (h *Handler) Handle(req Request) (Response, error) {
	stocks, err := h.businessLogic.Stocks(req.Sku)

	if err != nil {
		return Response{}, errors.Wrap(err, "creation failed")
	}

	responseStocks := make([]ResponseStock, len(stocks))
	for idx, item := range stocks {
		responseStocks[idx] = ResponseStock(item)
	}

	return Response{
		Stocks: responseStocks,
	}, nil
}
