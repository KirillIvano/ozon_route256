package loms_client

import (
	"net/http"
	"route256/checkout/internal/domain"
	"route256/libs/jsonreqwrap"
)

type StocksRequestBody struct {
	Sku uint32 `json:"sku"`
}

type StockResponseItem struct {
	Count       uint64 `json:"count"`
	WarehouseID uint64 `json:"warehouseID"`
}

type StocksResponse struct {
	Stocks []StockResponseItem `json:"stocks"`
}

func (c *Client) Stocks(sku uint32) ([]domain.Stock, error) {
	reqClient := jsonreqwrap.NewClient[StocksRequestBody, StocksResponse](
		c.urlStocks,
		http.MethodPost,
	)
	requestData := StocksRequestBody{Sku: sku}

	response, err := reqClient.Run(requestData)

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
