package service

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
)

func TestCreatTransanctionValid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountMock := repository.NewAccountMock()
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := NewTransactionService(transactionRepository)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO transactions(id, account_id, operation_type, amount) VALUES(?,?,?,?)").
		WithArgs(
			repository.IDMock{},
			accountMock.ID,
			transactionMock.OperationType,
			transactionMock.Amount*-1,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("UPDATE accounts set credit_limit = ? where id = ?").
		WithArgs(
			accountMock.AvailableCreditLimit+transactionMock.Amount*-1,
			accountMock.ID,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	account, _ := model.BuildAccount(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	amountModel, _ := model.NewAmount(transactionMock.Amount)
	operationTypeModel, _ := model.NewOperationType(transactionMock.OperationType)

	transaction, _ := model.NewTransaction(account, operationTypeModel, amountModel)

	err := transactionService.CreateTransaction(transaction)

	if err != nil {
		t.Errorf("\nAn error '%s' was not expected", err.Error())
	}
}
