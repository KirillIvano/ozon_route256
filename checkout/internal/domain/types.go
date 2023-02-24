package domain

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

type Offer struct {
	CartItem
	Price uint32
	Name  string
}

type Model struct {
	lomsService    LomsService
	productService ProductService
}
