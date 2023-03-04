package loms_server_v1

import (
	"context"
	lomsV1 "route256/loms/pkg/loms_v1"

	"github.com/pkg/errors"
)

var (
	ErrStocksEmptySku = errors.New("empty sku")
)

func (impl *implementation) Stocks(ctx context.Context, req *lomsV1.StocksParams) (*lomsV1.StocksResponse, error) {
	stocks, err := impl.lomsDomain.Stocks(req.Sku)

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
