package schema

type CartItem struct {
	UserId int64 `db:"user_id"`
	Count  int32 `db:"count"`
	Sku    int64 `db:"sku"`
}
