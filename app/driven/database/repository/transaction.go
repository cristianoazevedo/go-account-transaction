package repository

import (
	"database/sql"

	"github.com/csazevedo/go-account-transaction/app/model"
)

type transactionRepository struct {
	dbAdapter *sql.DB
}

//TransactionRepository interface representing the transaction repository struct
type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) error
}

//NewTransactionRepository creates a new struct of transaction repository
func NewTransactionRepository(adapter *sql.DB) TransactionRepository {
	return &transactionRepository{dbAdapter: adapter}
}

//CreateTransaction create a transaction
//As a critical point, the concept of transaction is used.
//If there is any problem, no changes are made
func (repository *transactionRepository) CreateTransaction(transaction model.Transaction) (err error) {
	tx, err := repository.dbAdapter.Begin()

	if err != nil {
		return
	}

	_, err = tx.Exec(
		"INSERT INTO transactions(id, account_id, operation_type, amount) VALUES(?,?,?,?)",
		transaction.GetID().GetValue(),
		transaction.GetAccount().GetID().GetValue(),
		transaction.GetOperationType().GetValue(),
		transaction.GetAmountValueByOperationType(),
	)

	if err != nil {
		tx.Rollback()

		return
	}

	_, err = tx.Exec(
		"UPDATE accounts set credit_limit = ? where id = ?",
		transaction.GetAccount().GetAvailableCreditLimit().GetValue(),
		transaction.GetAccount().GetID().GetValue(),
	)

	if err != nil {
		tx.Rollback()

		return
	}

	err = tx.Commit()

	return
}
