package domain

type LomsService interface {
	Stocks(sku uint32) ([]Stock, error)
	CreateOrder(user uint64, items []CartItem) (orderId uint64, err error)
}

type ProductService interface {
	GetProduct(sku uint32) (ProductInfo, error)
}

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type CartItem struct {
	UserId uint32
	Sku    uint32
	Count  uint16
}

type ProductInfo struct {
	Price uint32
	Name  string
}

// offer - item with price and name
type Offer struct {
	CartItem
	Price uint32
	Name  string
}

type Model struct {
	lomsService    LomsService
	productService ProductService
}

func New(lomsService LomsService, productService ProductService) *Model {
	return &Model{
		lomsService:    lomsService,
		productService: productService,
	}
}
