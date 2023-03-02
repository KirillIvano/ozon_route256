package main

import (
	"log"
	"net/http"
	loms_client "route256/checkout/internal/clients/loms"
	products_client "route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/handlers"
	"route256/libs/jsonhandlerwrap"
)

const PORT = ":8080"

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init failed")
	}

	lomsClient := loms_client.New(config.ConfigData.Services.Loms)
	productClient := products_client.New(config.ConfigData.Services.Products, config.ConfigData.Token)

	businessLogic := domain.New(lomsClient, productClient)

	handlersRegistry := handlers.New(businessLogic)

	http.Handle("/addToCart", jsonhandlerwrap.New(handlersRegistry.AddToCart))
	http.Handle("/deleteFromCart", jsonhandlerwrap.New(handlersRegistry.DeleteFromCart))
	http.Handle("/purchase", jsonhandlerwrap.New(handlersRegistry.Purchase))
	http.Handle("/listCart", jsonhandlerwrap.New(handlersRegistry.ListCart))

	log.Println("listening http at", PORT)
	err = http.ListenAndServe(PORT, nil)
	log.Fatal("cannot listen http", err)
}
