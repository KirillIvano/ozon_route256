package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"route256/libs/kafka"
	"route256/libs/logger"
	"route256/libs/tracing"
	"route256/libs/universal_server"
	"route256/loms/internal/config"
	"route256/loms/internal/domain"
	"route256/loms/internal/loms_server"
	orderSender "route256/loms/internal/order_sender"
	"route256/loms/internal/repository"
	statusActualizer "route256/loms/internal/status_actualizer"
	loms "route256/loms/pkg/loms_service"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	metricAddr = flag.String("addr", ":9081", "the port to listen")
	develMode  = flag.Bool("devel", false, "development mode")
)

func main() {
	ctx := context.Background()
	ctx, stopSignalListen := signal.NotifyContext(ctx, os.Interrupt)
	defer stopSignalListen()

	logger.Init(*develMode)
	tracing.Init("loms")

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		logger.Fatal("config init failed", zap.Error(err))
	}

	conn, err := pgxpool.New(ctx, config.ConfigData.Services.Database)
	if err != nil {
		logger.Fatal("pool connection error", zap.Error(err))
	}
	defer conn.Close()

	logger.Info("database connected successfully")

	producer, err := kafka.NewSyncProducer(config.ConfigData.Brokers)
	if err != nil {
		logger.Fatal("kafka connection failed", zap.Error(err))
	}
	orderSender := orderSender.NewOrderSender(producer, config.ConfigData.OrderTopic)

	repository := repository.Connect(conn)
	domain := domain.New(repository, orderSender)

	enf := statusActualizer.New(ctx, domain)

	enf.Start()
	defer enf.Close()

	server := universal_server.New("loms", config.ConfigData.Port, *metricAddr)

	loms.RegisterLomsServer(server.GetServerRegistrar(), loms_server.New(domain))

	server.Listen(ctx)
}
