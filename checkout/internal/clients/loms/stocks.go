package loms_client

import (
	"context"
	"route256/checkout/internal/domain"
	lomsService "route256/loms/pkg/loms_service"
)

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]domain.Stock, error) {
	requestData := lomsService.StocksParams{Sku: sku}

	response, err := c.client.Stocks(ctx, &requestData)

	if err != nil {
		return nil, err
	}

	stocks := make([]domain.Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, domain.Stock{
			WarehouseID: int64(stock.WarehouseID),
			Count:       stock.Count,
		})
	}

	return stocks, nil
}
