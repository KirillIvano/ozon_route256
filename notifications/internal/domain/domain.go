package domain

import (
	"route256/libs/logger"

	"go.uber.org/zap"
)

type notificationDomain struct{}

func (d *notificationDomain) LogOrderStatus(orderId int64, status string) error {
	logger.Info("order", zap.Int64("orderId", orderId), zap.String("status", status))

	return nil
}

func NewNotificationDomain() *notificationDomain {
	return &notificationDomain{}
}
