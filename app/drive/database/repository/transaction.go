package repository

import (
	"context"
	"database/sql"

	"github.com/csazevedo/go-account-transaction/app/model"
)

type transactionRepository struct {
	dbAdapter *sql.DB
}

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) error
}

func NewTransactionRepository(adapter *sql.DB) *transactionRepository {
	return &transactionRepository{dbAdapter: adapter}
}

func (repository *transactionRepository) CreateTransaction(transaction model.Transaction) error {
	ctx := context.Background()

	dbTransaction, err := repository.dbAdapter.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	query := "INSERT INTO transactions(id, account_id, operation_type, amount) VALUES(?,?,?,?)"

	insert, err := repository.dbAdapter.Prepare(query)

	if err != nil {
		return err
	}

	_, err = insert.ExecContext(
		ctx,
		transaction.GetId().GetValue(),
		transaction.GetAccount().GetId().GetValue(),
		transaction.GetOperationType().GetValue(),
		transaction.GetAmountValueByOperationType(),
	)

	if err != nil {
		dbTransaction.Rollback()

		return err
	}

	err = dbTransaction.Commit()

	if err != nil {
		return err
	}

	return nil
}
