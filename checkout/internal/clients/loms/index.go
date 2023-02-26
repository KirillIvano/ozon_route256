package loms_client

type Client struct {
	urlStocks      string
	urlCreateOrder string
}

func New(urlOrigin string) *Client {
	return &Client{
		urlStocks:      urlOrigin + "/stocks",
		urlCreateOrder: urlOrigin + "/createOrder",
	}
}
