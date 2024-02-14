package repository

import (
	"context"
)

// holds transaction specific methods
type RepositoryTransaction interface {
	// return a transaction
	BeginTx(ctx context.Context) (Transaction, error)
	HandleTransaction(ctx context.Context, tx Transaction, incomingErr error) error
}

type Transaction interface {
	Commit() error
	Rollback() error
}
