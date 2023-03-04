package loms_server_v1

import (
	"route256/loms/internal/domain"
	lomsV1 "route256/loms/pkg/loms_v1"
)

type implementation struct {
	lomsV1.UnimplementedLomsV1Server

	lomsDomain *domain.LomsDomain
}

var _ lomsV1.LomsV1Server = &implementation{}

func New(domain *domain.LomsDomain) *implementation {
	return &implementation{
		lomsDomain: domain,
	}
}
