package domain

func (m *Model) ListOrder(orderId int64) (OrderInfo, error) {

	return OrderInfo{
		Items: []OrderItem{
			{
				Sku:   1,
				Count: 10,
			},
		},
		User:   1,
		Status: OrderStatusAwaitingPayment,
	}, nil
}
