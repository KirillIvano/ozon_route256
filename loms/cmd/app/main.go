package main

import (
	"log"
	"net/http"
	"route256/libs/jsonhandlerwrap"
	"route256/loms/internal/config"
	cancelorder "route256/loms/internal/controllers/cancel_order"
	createorder "route256/loms/internal/controllers/create_order"
	listorder "route256/loms/internal/controllers/list_order"
	orderpayed "route256/loms/internal/controllers/order_payed"
	"route256/loms/internal/controllers/stocks"
	"route256/loms/internal/domain"
)

const PORT = ":8081"

func main() {
	config.Init()

	domain := domain.New()

	createOrderHandler := createorder.New(domain)
	listOrderHandler := listorder.New(domain)
	orderPayedHandler := orderpayed.New(domain)
	cancelOrderHandler := cancelorder.New(domain)
	stocksHandler := stocks.New(domain)

	http.Handle("/createOrder", jsonhandlerwrap.New(createOrderHandler.Handle))
	http.Handle("/listOrder", jsonhandlerwrap.New(listOrderHandler.Handle))
	http.Handle("/orderPayed", jsonhandlerwrap.New(orderPayedHandler.Handle))
	http.Handle("/cancelOrder", jsonhandlerwrap.New(cancelOrderHandler.Handle))
	http.Handle("/stocks", jsonhandlerwrap.New(stocksHandler.Handle))

	log.Println("listening http at", PORT)
	err := http.ListenAndServe(PORT, nil)
	log.Fatal("cannot listen http", err)
}
