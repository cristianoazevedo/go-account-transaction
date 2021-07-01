package repository

import (
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/google/uuid"
)

func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

type AccountMock struct {
	Id             string
	CreatedAt      string
	DocumentNumber string
}

func NewAccountMock() *AccountMock {
	return &AccountMock{
		Id:             uuid.New().String(),
		CreatedAt:      "2021-01-01 00:00:00",
		DocumentNumber: "03393983024",
	}
}

func TestCreateAccountValid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO accounts(id, document_number) VALUES(?, ?)").WithArgs(accountMock.Id, accountMock.DocumentNumber).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	account, _ := model.BuildAccount(accountMock.Id, accountMock.DocumentNumber, accountMock.CreatedAt)

	err := repository.CreateAccount(account)

	if err != nil {
		t.Errorf("\nError '%s' was not expected", err)
	}
}

func TestShouldRollbackCreateAccountOnFailure(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO accounts(id, document_number) VALUES(?, ?)").WithArgs(accountMock.Id, accountMock.DocumentNumber).WillReturnError(errors.New("timeout"))
	mock.ExpectRollback()

	account, _ := model.BuildAccount(accountMock.Id, accountMock.DocumentNumber, accountMock.CreatedAt)

	err := repository.CreateAccount(account)

	if err == nil {
		t.Errorf("\nError '%s' was expected", err)
	}
}

func TestShouldErroCreateAccountOnBeginDataBaseTransaction(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	mock.ExpectBegin().WillReturnError(errors.New("error"))

	account, _ := model.BuildAccount(accountMock.Id, accountMock.DocumentNumber, accountMock.CreatedAt)

	err := repository.CreateAccount(account)

	if err == nil {
		t.Errorf("\nError '%s' was expected", err)
	}
}

func TestFindByIDValid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.Id, accountMock.DocumentNumber, accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.Id).WillReturnRows(rows)

	idModel, _ := model.BuildId(accountMock.Id)

	account, err := repository.FindByID(idModel)

	if err != nil {
		t.Errorf("\nError '%s' was not expected", err)
	}

	if account.GetId().GetValue() != accountMock.Id {
		t.Errorf("\nInvalid ID value: '%v'", accountMock.Id)
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

	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.Id, accountMock.DocumentNumber, accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.DocumentNumber).WillReturnRows(rows)

	documentModel, _ := model.NewDocument(accountMock.DocumentNumber)

	account, err := repository.FindAccountByDocumentNumber(documentModel)

	if err != nil {
		t.Errorf("\nError '%s' was not expected", err)
	}

	if account.GetId().GetValue() != accountMock.Id {
		t.Errorf("\nInvalid ID value: '%v'", accountMock.Id)
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

	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{})

	mock.ExpectQuery(query).WithArgs(accountMock.Id).WillReturnRows(rows)

	idModel, _ := model.BuildId(accountMock.Id)

	account, err := repository.FindByID(idModel)

	if account != nil {
		t.Error("\nAccount was not expected")
	}

	if err != nil {
		t.Errorf("\nError '%s' was not expected", err)
	}
}

func TestFindWithAccountBuildInvalid(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.Id, "00000000001", accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.Id).WillReturnRows(rows)

	idModel, _ := model.BuildId(accountMock.Id)

	account, err := repository.FindByID(idModel)

	if account != nil {
		t.Error("\nAccount was not expected")
	}

	if err == nil {
		t.Errorf("\nError '%s' was expected", err)
	}
}

func TestFindWithDabaseError(t *testing.T) {
	db, mock := NewDBMock()
	repository := NewAccountRepository(db)
	accountMock := NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	mock.ExpectQuery(query).WithArgs(accountMock.Id).WillReturnError(errors.New("timeout"))

	idModel, _ := model.BuildId(accountMock.Id)

	account, err := repository.FindByID(idModel)

	if account != nil {
		t.Error("\nAccount was not expected")
	}

	if err == nil {
		t.Errorf("\nError '%s' was expected", err)
	}
}
