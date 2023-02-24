package main

import (
	"log"
	"net/http"
	loms_client "route256/checkout/internal/clients/loms"
	products_client "route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	addtocart "route256/checkout/internal/controllers/add_to_cart"
	deletfromcart "route256/checkout/internal/controllers/delete_from_cart"
	listcart "route256/checkout/internal/controllers/list_cart"
	"route256/checkout/internal/controllers/purchase"
	"route256/checkout/internal/domain"
	"route256/libs/jsonhandlerwrap"
)

const PORT = ":8080"

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init failed")
	}

	lomsClient := loms_client.New()
	productClient := products_client.New()

	businessLogic := domain.New(lomsClient, productClient)

	addToCartHandler := addtocart.New(businessLogic)
	deleteFromCartHandler := deletfromcart.New(businessLogic)
	purchaseHandler := purchase.New(businessLogic)
	listHandler := listcart.New(businessLogic)

	http.Handle("/addToCart", jsonhandlerwrap.New(addToCartHandler.Handle))
	http.Handle("/deleteFromCart", jsonhandlerwrap.New(deleteFromCartHandler.Handle))
	http.Handle("/purchase", jsonhandlerwrap.New(purchaseHandler.Handle))
	http.Handle("/listCart", jsonhandlerwrap.New(listHandler.Handle))

	log.Println("listening http at", PORT)
	err = http.ListenAndServe(PORT, nil)
	log.Fatal("cannot listen http", err)
}
