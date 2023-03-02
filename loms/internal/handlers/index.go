package handlers

import "route256/loms/internal/domain"

type LomsHandlersRegistry struct {
	domainLogic *domain.LomsDomain
}

func New(domainLogic *domain.LomsDomain) *LomsHandlersRegistry {
	return &LomsHandlersRegistry{
		domainLogic: domainLogic,
	}
}
