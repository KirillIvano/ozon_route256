package domain

type LomsService interface {
	Stocks(sku uint32) ([]Stock, error)
	CreateOrder(user uint64, items []CartItem) (orderId uint64, err error)
}

type ProductService interface {
	GetProduct(sku uint32) (ProductInfo, error)
}

type CheckoutDomain struct {
	lomsService    LomsService
	productService ProductService
}

func New(lomsService LomsService, productService ProductService) *CheckoutDomain {
	return &CheckoutDomain{
		lomsService:    lomsService,
		productService: productService,
	}
}
