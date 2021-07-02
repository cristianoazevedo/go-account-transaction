package usecases

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/service"
)

func TestCreateAccountWithDocumentInvalid(t *testing.T) {
	db, _ := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)

	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.CreateAccount("00000000001")

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError == nil {
		t.Error("\nDomain error was expected")
	}
}

func TestCreateAccountWithDocumentFound(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.DocumentNumber).WillReturnRows(rows)

	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.CreateAccount(accountMock.DocumentNumber)

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError == nil {
		t.Error("\nDomain error was expected")
	}
}

func TestCreateAccountWithFindDocumentWithError(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"

	mock.ExpectQuery(query).WithArgs(accountMock.DocumentNumber).WillReturnError(errors.New("timeout"))

	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.CreateAccount(accountMock.DocumentNumber)

	if infraError == nil {
		t.Error("\nInfra error was expected")
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was not expected", domainError.Error())
	}
}

func TestShouldRollbackCreateAccountOnFailure(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"
	rows := sqlmock.NewRows([]string{})
	mock.ExpectQuery(query).WithArgs(accountMock.DocumentNumber).WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO accounts(id, document_number) VALUES(?, ?)").
		WithArgs(repository.IDMock{}, accountMock.DocumentNumber).
		WillReturnError(errors.New("timeout"))
	mock.ExpectRollback()

	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.CreateAccount(accountMock.DocumentNumber)

	if infraError == nil {
		t.Error("\nInfra error was expected")
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was not expected", domainError.Error())
	}
}

func TestFindAccountFound(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.FindAccount(accountMock.ID)

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was not expected", domainError.Error())
	}
}

func TestFindAccountWithInfraError(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at"}).
		AddRow(accountMock.ID, "00000000001", accountMock.CreatedAt)

	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.FindAccount(accountMock.ID)

	if infraError == nil {
		t.Error("\nInfra error was expected")
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was not expected", domainError)
	}
}

func TestFindAccountWithDomainError(t *testing.T) {
	db, _ := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.FindAccount("99c49b65-cc11-487f-864d-55dbb6c90a6")

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError == nil {
		t.Error("\nDomain error was expected")
	}
}

func TestCreateAccountValid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)

	accountMock := repository.NewAccountMock()

	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"

	rows := sqlmock.NewRows([]string{})

	mock.ExpectQuery(query).WithArgs(accountMock.DocumentNumber).WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO accounts(id, document_number) VALUES(?, ?)").
		WithArgs(repository.IDMock{}, accountMock.DocumentNumber).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	usecase := NewAccountUseCase(accountSerivce)

	_, infraError, domainError := usecase.CreateAccount(accountMock.DocumentNumber)

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was not expected", domainError.Error())
	}
}
