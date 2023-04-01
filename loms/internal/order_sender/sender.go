package order_sender

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type orderSender struct {
	producer sarama.SyncProducer
	topic    string
}

type Handler func(id string)

func NewOrderSender(producer sarama.SyncProducer, topic string) *orderSender {
	return &orderSender{
		producer: producer,
		topic:    topic,
	}
}

type OrderInfo struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type OrderSender interface {
	SendOrder(ctx context.Context, orderId int64, orderStatus string) error
}

func (s *orderSender) SendOrder(ctx context.Context, orderId int64, orderStatus string) error {
	bytes, err := json.Marshal(OrderInfo{
		Id:     orderId,
		Status: orderStatus,
	})

	if err != nil {
		return errors.Wrap(err, "failed to marshal message")
	}

	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(bytes),
		Key:       sarama.StringEncoder(fmt.Sprint(orderId)),
		Timestamp: time.Now(),
	}

	_, _, err = s.producer.SendMessage(msg)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	return nil
}
