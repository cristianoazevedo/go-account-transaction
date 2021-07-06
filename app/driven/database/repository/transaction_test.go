package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csazevedo/go-account-transaction/app/model"
)

func TestCreateTransactionValid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewTransactionRepository(db)
	accountMock := NewAccountMock()
	transactionMock := NewTransactionMock()

	accountModel, _ := model.BuildAccount(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	operationTypeModel, _ := model.NewOperationType(transactionMock.OperationType)
	amountModel, _ := model.NewAmount(transactionMock.Amount)

	transaction, _ := model.NewTransaction(accountModel, operationTypeModel, amountModel)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO transactions(id, account_id, operation_type, amount) VALUES(?,?,?,?)").
		WithArgs(
			IDMock{},
			accountMock.ID,
			transactionMock.OperationType,
			transactionMock.Amount*-1,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("UPDATE accounts set credit_limit = ? where id = ?").
		WithArgs(
			transaction.GetAccount().GetAvailableCreditLimit().GetValue(),
			accountMock.ID,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	err := repository.CreateTransaction(transaction)

	if err != nil {
		t.Errorf("\nAn error: '%s' was not expected", err.Error())
	}
}

func TestShouldRollbackCreateTransactionOnFailure(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewTransactionRepository(db)
	accountMock := NewAccountMock()
	transactionMock := NewTransactionMock()

	accountModel, _ := model.BuildAccount(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	operationTypeModel, _ := model.NewOperationType(transactionMock.OperationType)
	amountModel, _ := model.NewAmount(transactionMock.Amount)

	transaction, _ := model.NewTransaction(accountModel, operationTypeModel, amountModel)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO transactions(id, account_id, operation_type, amount) VALUES(?,?,?,?)").
		WithArgs(
			IDMock{},
			accountMock.ID,
			transactionMock.OperationType,
			transactionMock.Amount*-1,
		).
		WillReturnError(errors.New("timeout"))
	mock.ExpectRollback()

	err := repository.CreateTransaction(transaction)

	if err == nil {
		t.Error("\nAn error was expected")
	}
}

func TestShouldErroCreateTransactionOnBeginDataBaseTransaction(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewTransactionRepository(db)
	accountMock := NewAccountMock()
	transactionMock := NewTransactionMock()

	accountModel, _ := model.BuildAccount(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	operationTypeModel, _ := model.NewOperationType(transactionMock.OperationType)
	amountModel, _ := model.NewAmount(transactionMock.Amount)

	transaction, _ := model.NewTransaction(accountModel, operationTypeModel, amountModel)

	mock.ExpectBegin().WillReturnError(errors.New("error"))

	err := repository.CreateTransaction(transaction)

	if err == nil {
		t.Error("\nAn error was expected")
	}
}
