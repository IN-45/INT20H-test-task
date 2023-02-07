package db

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type TransactionManager struct {
	db *bun.DB
}

func NewTransactionManager(db *bun.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (m *TransactionManager) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	transaction, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if err := tFunc(injectTransaction(ctx, &transaction)); err != nil {
		if rollbackErr := transaction.Rollback(); rollbackErr != nil {
			return rollbackErr
		}

		return err
	}

	return transaction.Commit()
}

func injectTransaction(ctx context.Context, tx *bun.Tx) context.Context {
	return context.WithValue(ctx, transactionKey{}, tx)
}
