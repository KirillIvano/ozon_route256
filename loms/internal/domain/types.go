package domain

type UserId = int64

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
	User   UserId
	Items  []OrderItem
}

type Reservation struct {
	WarehouseId int64
	OrderId     int64
	Sku         uint32
	Count       int32
}
