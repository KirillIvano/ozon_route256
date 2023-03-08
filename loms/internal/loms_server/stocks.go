package loms_server

import (
	"context"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/pkg/errors"
)

var (
	ErrStocksEmptySku = errors.New("empty sku")
)

func ValidateStocksParams(r *lomsService.StocksParams) error {
	if r.Sku == 0 {
		return ErrStocksEmptySku
	}

	return nil
}

func (impl *implementation) Stocks(ctx context.Context, params *lomsService.StocksParams) (*lomsService.StocksResponse, error) {
	if err := ValidateStocksParams(params); err != nil {
		return nil, err
	}

	stocks, err := impl.lomsDomain.Stocks(params.Sku)

	if err != nil {
		return nil, errors.Wrap(err, "creation failed")
	}

	responseStocks := make([]*lomsService.StocksResponseItem, len(stocks))
	for idx, item := range stocks {
		responseStocks[idx] = &lomsService.StocksResponseItem{
			WarehouseID: item.WarehouseID,
			Count:       item.Count,
		}
	}

	return &lomsService.StocksResponse{
		Stocks: responseStocks,
	}, nil
}
