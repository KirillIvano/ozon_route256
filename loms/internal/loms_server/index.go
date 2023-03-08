package loms_server

import (
	"route256/loms/internal/domain"
	lomsService "route256/loms/pkg/loms_service"
)

type implementation struct {
	lomsService.UnimplementedLomsServer

	lomsDomain *domain.LomsDomain
}

var _ lomsService.LomsServer = &implementation{}

func New(domain *domain.LomsDomain) *implementation {
	return &implementation{
		lomsDomain: domain,
	}
}
