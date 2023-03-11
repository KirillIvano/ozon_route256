package repository

import (
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/transactor"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	ConnManager *transactor.TransactionManager
}

var _ domain.LomsRepository = (*repository)(nil)

func Connect(conn *pgx.Conn) repository {
	return repository{
		ConnManager: transactor.New(conn),
	}
}
