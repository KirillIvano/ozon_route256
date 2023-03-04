package loms_server_v1

import (
	"context"
	lomsV1 "route256/loms/pkg/loms_v1"

	"github.com/pkg/errors"
)

var (
	ErrStocksEmptySku = errors.New("empty sku")
)

func ValidateStocksParams(r *lomsV1.StocksParams) error {
	if r.Sku == 0 {
		return ErrStocksEmptySku
	}

	return nil
}

func (impl *implementation) Stocks(ctx context.Context, params *lomsV1.StocksParams) (*lomsV1.StocksResponse, error) {
	if err := ValidateStocksParams(params); err != nil {
		return nil, err
	}

	stocks, err := impl.lomsDomain.Stocks(params.Sku)

	if err != nil {
		return nil, errors.Wrap(err, "creation failed")
	}

	responseStocks := make([]*lomsV1.StocksResponseItem, len(stocks))
	for idx, item := range stocks {
		responseStocks[idx] = &lomsV1.StocksResponseItem{
			WarehouseID: item.WarehouseID,
			Count:       item.Count,
		}
	}

	return &lomsV1.StocksResponse{
		Stocks: responseStocks,
	}, nil
}
