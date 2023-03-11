package domain

import "context"

type LomsService interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, user int64, items []CartItem) (orderId int64, err error)
}

type ProductService interface {
	GetProduct(ctx context.Context, sku uint32) (ProductInfo, error)
}

type CheckoutRepository interface {
	DeleteCart(ctx context.Context, userId int64) error
	GetCartItems(ctx context.Context, userId int64) ([]CartItem, error)
	DeleteItem(ctx context.Context, userId int64, sku uint32, count uint32) error
	AddToCart(ctx context.Context, userId int64, sku uint32, count uint32) error
}
type CheckoutDomain struct {
	lomsService    LomsService
	productService ProductService
	repository     CheckoutRepository
}

func New(lomsService LomsService, productService ProductService, repository CheckoutRepository) *CheckoutDomain {
	return &CheckoutDomain{
		lomsService:    lomsService,
		productService: productService,
		repository:     repository,
	}
}
