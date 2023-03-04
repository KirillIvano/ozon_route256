package checkout_server_v1

import (
	"route256/checkout/internal/domain"
	checkoutV1 "route256/checkout/pkg/checkout_v1"
)

type implementation struct {
	checkoutV1.UnimplementedCheckoutV1Server

	checkoutDomain *domain.CheckoutDomain
}

var _ checkoutV1.CheckoutV1Server = &implementation{}

func New(domain *domain.CheckoutDomain) *implementation {
	return &implementation{
		checkoutDomain: domain,
	}
}
