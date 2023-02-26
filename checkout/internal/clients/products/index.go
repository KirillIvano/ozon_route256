package products_client

type Client struct {
	token         string
	urlGetProduct string
}

func New(urlOrigin string, token string) *Client {
	return &Client{
		token:         token,
		urlGetProduct: urlOrigin + "/get_product",
	}
}
