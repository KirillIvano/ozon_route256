package controllers

import "route256/loms/internal/domain"

type LomsHandlersRegistry struct {
	domainLogic *domain.LomsDomain
}

func NewLomsHandlersRegistry(domainLogic *domain.LomsDomain) *LomsHandlersRegistry {
	return &LomsHandlersRegistry{
		domainLogic: domainLogic,
	}
}
