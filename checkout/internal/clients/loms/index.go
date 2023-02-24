package loms_client

import "route256/checkout/internal/config"

type Client struct {
	urlStocks      string
	urlCreateOrder string
}

func New() *Client {
	return &Client{
		urlStocks:      config.ConfigData.Services.Loms + "/stocks",
		urlCreateOrder: config.ConfigData.Services.Loms + "/createOrder",
	}
}
