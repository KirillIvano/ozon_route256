package domain

import "log"

type notificationDomain struct{}

func (d *notificationDomain) LogOrderStatus(orderId int64, status string) error {
	log.Printf("order %d: %s", orderId, status)

	return nil
}

func NewNotificationDomain() *notificationDomain {
	return &notificationDomain{}
}
