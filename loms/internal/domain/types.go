package domain

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type OrderItem struct {
	Sku   uint32
	Count uint32
}

type OrderInfo struct {
	Status string // (new | awaiting payment | failed | payed | cancelled)
	User   int64
	Items  []OrderItem
}

var (
	OrderStatusNew             = "new"
	OrderStatusAwaitingPayment = "awaiting payment"
	OrderStatusFailed          = "failed"
	OrderStatusPayed           = "payed"
	OrderStatusCancelled       = "cancelled"
)
