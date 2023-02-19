package products_client

import (
	"net/http"
	"route256/checkout/internal/domain"
	"route256/libs/jsonreqwrap"
)

type GetProductRequest struct {
	Sku   uint32 `json:"sku"`
	Token string `json:"token"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *Client) GetProduct(sku uint32) (domain.ProductInfo, error) {
	reqClient := jsonreqwrap.NewClient[GetProductRequest, GetProductResponse](
		c.urlGetProduct,
		http.MethodPost,
	)
	requestData := GetProductRequest{
		Sku:   sku,
		Token: c.token,
	}

	response, err := reqClient.Run(requestData)

	if err != nil {
		return domain.ProductInfo{}, err
	}

	return domain.ProductInfo{
		Price: response.Price,
		Name:  response.Name,
	}, nil
}
