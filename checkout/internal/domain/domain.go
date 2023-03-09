package domain

import "context"

type LomsService interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, user int64, items []CartItem) (orderId int64, err error)
}

type ProductService interface {
	GetProduct(ctx context.Context, sku uint32) (ProductInfo, error)
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
