package domain

func (m *Model) Stocks(sku uint32) ([]Stock, error) {

	return []Stock{
		{
			WarehouseID: 1,
			Count:       1,
		},
		{
			WarehouseID: 1,
			Count:       2,
		},
	}, nil
}
