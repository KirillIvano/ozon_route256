package main

import (
	"log"
	"net/http"
	"route256/libs/jsonhandlerwrap"
	"route256/loms/internal/controllers"
	"route256/loms/internal/domain"
)

const PORT = ":8081"

func main() {
	domain := domain.New()

	handlersRegistry := controllers.NewLomsHandlersRegistry(domain)

	http.Handle("/createOrder", jsonhandlerwrap.New(handlersRegistry.HandleCreateOrder))
	http.Handle("/listOrder", jsonhandlerwrap.New(handlersRegistry.HandleListOrder))
	http.Handle("/orderPayed", jsonhandlerwrap.New(handlersRegistry.HandleOrderPayed))
	http.Handle("/cancelOrder", jsonhandlerwrap.New(handlersRegistry.HandleCancelOrder))
	http.Handle("/stocks", jsonhandlerwrap.New(handlersRegistry.HandleStocks))

	log.Println("listening http at", PORT)
	err := http.ListenAndServe(PORT, nil)
	log.Fatal("cannot listen http", err)
}
