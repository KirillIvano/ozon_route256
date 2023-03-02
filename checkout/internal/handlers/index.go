package handlers

import "route256/checkout/internal/domain"

type CheckoutHandlersRegistry struct {
	domainLogic *domain.CheckoutDomain
}

func New(businessLogic *domain.CheckoutDomain) *CheckoutHandlersRegistry {
	return &CheckoutHandlersRegistry{
		domainLogic: businessLogic,
	}
}
