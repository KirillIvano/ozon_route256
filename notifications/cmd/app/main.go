package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"route256/libs/logger"
	"route256/notifications/internal/config"
	"route256/notifications/internal/domain"
	"route256/notifications/internal/order_consumer"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

func main() {
	logger.Init(false)

	err := config.Init()
	if err != nil {
		logger.Fatal("config init failed", zap.Error(err))
	}

	keepRunning := true

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = sarama.MaxVersion
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	kafkaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}

	domain := domain.NewNotificationDomain()
	consumer := order_consumer.NewConsumerGroup(domain)

	const groupName = "group-+v"

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(config.ConfigData.Brokers, groupName, kafkaConfig)
	if err != nil {
		logger.Fatal("Error creating consumer group client", zap.Error(err))
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{config.ConfigData.OrderTopic}, &consumer); err != nil {
				logger.Fatal("Error from consumer", zap.Error(err))

			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-consumer.Ready()
	logger.Info("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			logger.Info("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			logger.Info("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}

	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		logger.Fatal("Error closing client", zap.Error(err))
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		logger.Info("Resuming consumption")
	} else {
		client.PauseAll()
		logger.Info("Pausing consumption")
	}

	*isPaused = !*isPaused
}
