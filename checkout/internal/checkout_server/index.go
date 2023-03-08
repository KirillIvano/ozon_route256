package checkout_server

import (
	"route256/checkout/internal/domain"
	checkoutService "route256/checkout/pkg/checkout_service"
)

type implementation struct {
	checkoutService.UnimplementedCheckoutServer

	checkoutDomain *domain.CheckoutDomain
}

var _ checkoutService.CheckoutServer = &implementation{}

func New(domain *domain.CheckoutDomain) *implementation {
	return &implementation{
		checkoutDomain: domain,
	}
}
