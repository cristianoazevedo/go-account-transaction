package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csazevedo/go-account-transaction/app/model"
)

func TestCreateAccountValid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO accounts(id, document_number, credit_limit) VALUES(?, ?, ?)").
		WithArgs(IDMock{}, accountMock.DocumentNumber, accountMock.AvailableCreditLimit).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	document, _ := model.NewDocument(accountMock.DocumentNumber)
	creditLimit, _ := model.BuildAvailableCreditLimit(accountMock.AvailableCreditLimit)
	account := model.NewAccount(document, creditLimit)

	err := repository.CreateAccount(account)

	if err != nil {
		t.Errorf("\nAn error: '%s' was not expected", err.Error())
	}
}

func TestShouldRollbackCreateAccountOnFailure(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO accounts(id, document_number, credit_limit) VALUES(?, ?, ?)").
		WithArgs(IDMock{}, accountMock.DocumentNumber, accountMock.AvailableCreditLimit).
		WillReturnError(errors.New("timeout"))
	mock.ExpectRollback()

	document, _ := model.NewDocument(accountMock.DocumentNumber)
	creditLimit, _ := model.BuildAvailableCreditLimit(accountMock.AvailableCreditLimit)
	account := model.NewAccount(document, creditLimit)

	err := repository.CreateAccount(account)

	if err == nil {
		t.Error("\nAn error was expected")
	}
}

func TestShouldErroCreateAccountOnBeginDataBaseTransaction(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	mock.ExpectBegin().WillReturnError(errors.New("error"))

	document, _ := model.NewDocument(accountMock.DocumentNumber)
	creditLimit, _ := model.BuildAvailableCreditLimit(accountMock.AvailableCreditLimit)
	account := model.NewAccount(document, creditLimit)

	err := repository.CreateAccount(account)

	if err == nil {
		t.Error("\nAn error was expected")
	}
}

func TestFindByIDValid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at", "credit_limit"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)

	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	idModel, _ := model.BuildID(accountMock.ID)

	account, err := repository.FindByID(idModel)

	if err != nil {
		t.Errorf("\nAn error: '%s' was not expected", err.Error())
	}

	if account.GetID().GetValue() != accountMock.ID {
		t.Errorf("\nInvalid ID value: '%v'", accountMock.ID)
	}

	if account.GetDocument().GetValue() != accountMock.DocumentNumber {
		t.Errorf("\nInvalid document value: '%v'", accountMock.DocumentNumber)
	}

	if account.GetCreatedAt().GetValue() != accountMock.CreatedAt {
		t.Errorf("\nInvalid date value: '%v'", accountMock.CreatedAt)
	}
}

func TestFindByDocumentValid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where document_number = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at", "credit_limit"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)

	mock.ExpectQuery(query).WithArgs(accountMock.DocumentNumber).WillReturnRows(rows)

	documentModel, _ := model.NewDocument(accountMock.DocumentNumber)

	account, err := repository.FindAccountByDocumentNumber(documentModel)

	if err != nil {
		t.Errorf("\nAn error: '%s' was not expected", err.Error())
	}

	if account.GetID().GetValue() != accountMock.ID {
		t.Errorf("\nInvalid ID value: '%v'", accountMock.ID)
	}

	if account.GetDocument().GetValue() != accountMock.DocumentNumber {
		t.Errorf("\nInvalid document value: '%v'", accountMock.DocumentNumber)
	}

	if account.GetCreatedAt().GetValue() != accountMock.CreatedAt {
		t.Errorf("\nInvalid date value: '%v'", accountMock.CreatedAt)
	}
}

func TestFindWithoutResult(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{})

	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	idModel, _ := model.BuildID(accountMock.ID)

	account, err := repository.FindByID(idModel)

	if account != nil {
		t.Error("\nAccount was not expected")
	}

	if err != nil {
		t.Errorf("\nAn error: '%s' was not expected", err.Error())
	}
}

func TestFindWithAccountBuildInvalid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at", "credit_limit"}).
		AddRow(accountMock.ID, "00000000001", accountMock.CreatedAt, accountMock.AvailableCreditLimit)

	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	idModel, _ := model.BuildID(accountMock.ID)

	account, err := repository.FindByID(idModel)

	if account != nil {
		t.Error("\nAccount was not expected")
	}

	if err == nil {
		t.Error("\nAn error was expected")
	}
}

func TestFindWithDatabaseError(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"

	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnError(errors.New("timeout"))

	idModel, _ := model.BuildID(accountMock.ID)

	account, err := repository.FindByID(idModel)

	if account != nil {
		t.Error("\nAccount was not expected")
	}

	if err == nil {
		t.Error("\nAn Error was expected")
	}
}
