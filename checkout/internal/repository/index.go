package repository

import (
	"route256/checkout/internal/domain"
	"route256/libs/transactor"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	connManager *transactor.TransactionManager
}

var _ domain.CheckoutRepository = (*repository)(nil)

func New(conn *pgxpool.Pool) repository {
	return repository{
		connManager: transactor.New(conn),
	}
}
