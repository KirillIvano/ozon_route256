package main

import (
	"log"
	"net/http"
	loms_client "route256/checkout/internal/clients/loms"
	products_client "route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/controllers"
	"route256/checkout/internal/domain"
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
	controllersRegistry := controllers.NewCheckoutHandlersRegistry(businessLogic)

	http.Handle("/addToCart", jsonhandlerwrap.New(controllersRegistry.HandleAddToCart))
	http.Handle("/deleteFromCart", jsonhandlerwrap.New(controllersRegistry.HandleDeleteFromCart))
	http.Handle("/purchase", jsonhandlerwrap.New(controllersRegistry.HandlePurchase))
	http.Handle("/listCart", jsonhandlerwrap.New(controllersRegistry.HandleListCart))

	log.Println("listening http at", PORT)
	err = http.ListenAndServe(PORT, nil)
	log.Fatal("cannot listen http", err)
}
