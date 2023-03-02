package main

import (
	"log"
	"net/http"
	"route256/libs/jsonhandlerwrap"
	"route256/loms/internal/domain"
	"route256/loms/internal/handlers"
)

const PORT = ":8081"

func main() {
	domain := domain.New()

	handlersRegistry := handlers.New(domain)

	http.Handle("/createOrder", jsonhandlerwrap.New(handlersRegistry.CreateOrder))
	http.Handle("/listOrder", jsonhandlerwrap.New(handlersRegistry.ListOrder))
	http.Handle("/orderPayed", jsonhandlerwrap.New(handlersRegistry.OrderPayed))
	http.Handle("/cancelOrder", jsonhandlerwrap.New(handlersRegistry.CancelOrder))
	http.Handle("/stocks", jsonhandlerwrap.New(handlersRegistry.Stocks))

	log.Println("listening http at", PORT)
	err := http.ListenAndServe(PORT, nil)
	log.Fatal("cannot listen http", err)
}
