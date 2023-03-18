package repository

import (
	"route256/libs/transactor"
	"route256/loms/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	connManager *transactor.TransactionManager
}

var _ domain.LomsRepository = (*repository)(nil)

func Connect(conn *pgxpool.Pool) repository {
	return repository{
		connManager: transactor.New(conn),
	}
}
