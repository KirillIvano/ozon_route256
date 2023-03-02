package handlers

import (
	"github.com/pkg/errors"
)

type StocksRequest struct {
	Sku uint32 `json:"sku"`
}

var (
	ErrStocksEmptySku = errors.New("empty sku")
)

func (r StocksRequest) Validate() error {
	if r.Sku == 0 {
		return ErrStocksEmptySku
	}

	return nil
}

type ResponseStock = struct {
	WarehouseID int64  `json:"warehouseId"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []ResponseStock `json:"stocks"`
}

func (h *LomsHandlersRegistry) Stocks(req StocksRequest) (StocksResponse, error) {
	stocks, err := h.domainLogic.Stocks(req.Sku)

	if err != nil {
		return StocksResponse{}, errors.Wrap(err, "creation failed")
	}

	responseStocks := make([]ResponseStock, len(stocks))
	for idx, item := range stocks {
		responseStocks[idx] = ResponseStock{
			WarehouseID: item.WarehouseID,
			Count:       item.Count,
		}
	}

	return StocksResponse{
		Stocks: responseStocks,
	}, nil
}
