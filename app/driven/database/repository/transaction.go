package repository

import (
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

func (repository *transactionRepository) CreateTransaction(transaction model.Transaction) (err error) {
	tx, err := repository.dbAdapter.Begin()

	if err != nil {
		return
	}

	_, err = tx.Exec(
		"INSERT INTO transactions(id, account_id, operation_type, amount) VALUES(?,?,?,?)",
		transaction.GetId().GetValue(),
		transaction.GetAccount().GetId().GetValue(),
		transaction.GetOperationType().GetValue(),
		transaction.GetAmountValueByOperationType(),
	)

	if err != nil {
		tx.Rollback()

		return
	}

	err = tx.Commit()

	return
}
