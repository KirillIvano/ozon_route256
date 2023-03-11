package repository

import (
	"route256/libs/transactor"
	"route256/loms/internal/domain"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	connManager *transactor.TransactionManager
}

var _ domain.LomsRepository = (*repository)(nil)

func Connect(conn *pgx.Conn) repository {
	return repository{
		connManager: transactor.New(conn),
	}
}
