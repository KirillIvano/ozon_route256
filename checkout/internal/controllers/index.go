package controllers

import "route256/checkout/internal/domain"

type CheckoutHandlersRegistry struct {
	domainLogic *domain.CheckoutDomain
}

func NewCheckoutHandlersRegistry(businessLogic *domain.CheckoutDomain) *CheckoutHandlersRegistry {
	return &CheckoutHandlersRegistry{
		domainLogic: businessLogic,
	}
}
