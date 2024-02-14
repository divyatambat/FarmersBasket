package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "github.com/jackc/pgx/v4"
)

type BaseRepository struct {
	DB *sqlx.DB
}

type BaseTransaction struct {
	tx *sqlx.Tx
}

func (repo *BaseRepository) TimeNow() time.Time {
	return time.Now()
}

func (repo *BaseRepository) BeginTx(ctx context.Context) (Transaction, error) {
	tx, err := repo.DB.BeginTxx(ctx, nil)
	if err != nil {
		fmt.Printf("error occured while initiating database transaction: %v", err)
		return nil, err
	}

	return &BaseTransaction{
		tx: tx,
	}, nil
}

func (repo *BaseRepository) HandleTransaction(ctx context.Context, tx Transaction, incomingErr error) (err error) {
	if incomingErr != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Printf("error occured while rollback database transaction: %v", err)
			return
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("error occured while commit database transaction: %v", err)
		return
	}
	return
}

func (bt *BaseTransaction) Commit() error {
	return bt.tx.Commit()
}

func (bt *BaseTransaction) Rollback() error {
	return bt.tx.Rollback()
}

func (bt *BaseRepository) initiateQueryExecutor(tx Transaction) (QueryExecutor interface{}) {
	if tx != nil {
		bt := tx.(*BaseTransaction)
		return bt.tx
	}
	return bt.DB
}

type Transaction interface {
	Commit() error
	Rollback() error
}
