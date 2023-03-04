package transformers

import (
	"route256/loms/internal/domain"
	lomsV1 "route256/loms/pkg/loms_v1"
)

var (
	EnumToOrderStatus = map[lomsV1.OrderStatus]string{
		0: domain.OrderStatusUnspecified,
		1: domain.OrderStatusNew,
		2: domain.OrderStatusAwaitingPayment,
		3: domain.OrderStatusFailed,
		4: domain.OrderStatusPayed,
		5: domain.OrderStatusCancelled,
	}
	OrderStatusToEnum = map[string]lomsV1.OrderStatus{
		domain.OrderStatusUnspecified:     0,
		domain.OrderStatusNew:             1,
		domain.OrderStatusAwaitingPayment: 2,
		domain.OrderStatusFailed:          3,
		domain.OrderStatusPayed:           4,
		domain.OrderStatusCancelled:       5,
	}
)

func TransformOrderStatusToDomain(status lomsV1.OrderStatus) string {
	val, ok := EnumToOrderStatus[status]

	if !ok {
		return EnumToOrderStatus[0]
	}

	return val
}

func TransformDomainOrderStatusToApi(status string) lomsV1.OrderStatus {
	return OrderStatusToEnum[status]
}
