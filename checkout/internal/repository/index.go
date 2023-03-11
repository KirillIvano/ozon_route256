package repository

import (
	"route256/checkout/internal/domain"
	"route256/libs/transactor"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	connManager *transactor.TransactionManager
}

var _ domain.CheckoutRepository = (*repository)(nil)

func New(conn *pgx.Conn) repository {
	return repository{
		connManager: transactor.New(conn),
	}
}
