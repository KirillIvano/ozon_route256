package model

type Order struct {
	OrderId     int64  `db:"order_id"`
	UserId      int64  `db:"user_id"`
	OrderStatus string `db:"order_status"`
	CreatedAt   string `db:"created_at"`
}
