package products_client

import "route256/checkout/internal/config"

type Client struct {
	token         string
	urlGetProduct string
}

func New() *Client {
	return &Client{
		token:         config.ConfigData.Token,
		urlGetProduct: "http://route256.pavl.uk:8080/get_product",
	}
}
