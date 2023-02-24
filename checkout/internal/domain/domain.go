package domain

type LomsService interface {
	Stocks(sku uint32) ([]Stock, error)
	CreateOrder(user uint64, items []CartItem) (orderId uint64, err error)
}

type ProductService interface {
	GetProduct(sku uint32) (ProductInfo, error)
}

func New(lomsService LomsService, productService ProductService) *Model {
	return &Model{
		lomsService:    lomsService,
		productService: productService,
	}
}
