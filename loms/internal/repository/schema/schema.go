package schema

import "time"

type Order struct {
	OrderId     int64     `db:"order_id"`
	UserId      int64     `db:"user_id"`
	OrderStatus string    `db:"order_status"`
	CreatedAt   time.Time `db:"created_at"`
}

type WarehouseItem struct {
	WarehouseId int64 `db:"warehouse_id"`
	Sku         int64 `db:"sku"`
	Count       int32 `db:"count"`
}

type Reservation struct {
	WarehouseId int64 `db:"warehouse_id"`
	OrderId     int64 `db:"order_id"`
	Sku         int64 `db:"sku"`
	Count       int32 `db:"count"`
}

type OrderItem struct {
	OrderId int64 `db:"order_id"`
	Sku     int64 `db:"sku"`
	Count   int32 `db:"count"`
}
