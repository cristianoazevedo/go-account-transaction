package service

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
)

func TestCreateAccountValid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO accounts(id, document_number) VALUES(?, ?)").
		WithArgs(repository.IDMock{}, accountMock.DocumentNumber).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	account, _ := model.BuildAccount(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt)

	err := accountSerivce.CreateAccount(account)

	if err != nil {
		t.Errorf("\nAn error '%s' was not expected", err.Error())
	}
}

func TestFindAccountByIDValid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	idModel, _ := model.BuildID(accountMock.ID)

	_, err := accountSerivce.FindAccountByID(idModel)

	if err != nil {
		t.Errorf("\nAn error '%s' was not expected", err.Error())
	}
}

func TestFindAccountByDocumentNumberValid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.DocumentNumber).WillReturnRows(rows)

	documentModel, _ := model.NewDocument(accountMock.DocumentNumber)

	_, err := accountSerivce.FindAccountByDocumentNumber(documentModel)

	if err != nil {
		t.Errorf("\nAn error '%s' was not expected", err.Error())
	}
}
