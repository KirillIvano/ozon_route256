package order_consumer

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type NotificationDomain interface {
	LogOrderStatus(orderId int64, status string) error
}

type Consumer struct {
	ready  chan bool
	domain NotificationDomain
}

func NewConsumerGroup(domain NotificationDomain) Consumer {
	return Consumer{
		ready:  make(chan bool),
		domain: domain,
	}
}

func (consumer *Consumer) Ready() <-chan bool {
	return consumer.ready
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

type OrderInfo struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			var order OrderInfo
			err := json.Unmarshal(message.Value, &order)

			if err != nil {
				log.Printf("failed to parse order for key %s: %s\n", message.Key, err)
				session.MarkMessage(message, "")
				break
			}

			err = consumer.domain.LogOrderStatus(order.Id, order.Status)
			if err != nil {
				log.Println("failed to log status")
				break
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
