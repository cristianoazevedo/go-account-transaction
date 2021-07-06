package usecases

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/service"
)

func TestCreateTransactionValid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"
	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at", "credit_limit"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

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

	usecase := NewTransactionUseCase(transactionService, accountSerivce)

	_, infraError, domainError := usecase.CreateTransaction(accountMock.ID, transactionMock.OperationType, transactionMock.Amount)

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was not expected", domainError.Error())
	}
}

func TestCreateTransactionInvalid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"
	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at", "credit_limit"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO transactions(id, account_id, operation_type, amount) VALUES(?,?,?,?)").
		WithArgs(
			repository.IDMock{},
			accountMock.ID,
			transactionMock.OperationType,
			transactionMock.Amount*-1,
		).
		WillReturnError(errors.New("timeout"))
	mock.ExpectRollback()

	usecase := NewTransactionUseCase(transactionService, accountSerivce)

	_, infraError, domainError := usecase.CreateTransaction(accountMock.ID, transactionMock.OperationType, transactionMock.Amount)

	if infraError == nil {
		t.Error("\nInfra error was expected")
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was not expected", domainError.Error())
	}
}

func TestCreateTransactionWithAccountNotFound(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"
	rows := sqlmock.NewRows([]string{})
	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	usecase := NewTransactionUseCase(transactionService, accountSerivce)

	_, infraError, domainError := usecase.CreateTransaction(accountMock.ID, transactionMock.OperationType, transactionMock.Amount)

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError == nil {
		t.Error("\nDomain error was expected")
	}
}

func TestCreateTransactionWithFindAccountWithError(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"
	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnError(errors.New("timeout"))

	usecase := NewTransactionUseCase(transactionService, accountSerivce)

	_, infraError, domainError := usecase.CreateTransaction(accountMock.ID, transactionMock.OperationType, transactionMock.Amount)

	if infraError == nil {
		t.Error("\nInfra error was expected")
	}

	if domainError != nil {
		t.Errorf("\nDomain error: '%s' was expected", domainError.Error())
	}
}

func TestCreateTransactionWithAmountInvalid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"
	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at", "credit_limit"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	usecase := NewTransactionUseCase(transactionService, accountSerivce)

	_, infraError, domainError := usecase.CreateTransaction(accountMock.ID, transactionMock.OperationType, 0)

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError == nil {
		t.Error("\nDomain error was expected")
	}
}

func TestCreateTransactionWithOperationTypeInvalid(t *testing.T) {
	db, mock := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	accountMock := repository.NewAccountMock()
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	query := "SELECT id, document_number, created_at, credit_limit FROM accounts where id = ?"
	rows := sqlmock.NewRows([]string{"id", "document_number", "created_at", "credit_limit"}).
		AddRow(accountMock.ID, accountMock.DocumentNumber, accountMock.CreatedAt, accountMock.AvailableCreditLimit)
	mock.ExpectQuery(query).WithArgs(accountMock.ID).WillReturnRows(rows)

	usecase := NewTransactionUseCase(transactionService, accountSerivce)

	_, infraError, domainError := usecase.CreateTransaction(accountMock.ID, 5, transactionMock.Amount)

	if infraError != nil {
		t.Errorf("\nInfra error: '%s' was not expected", infraError.Error())
	}

	if domainError == nil {
		t.Error("\nDomain error was expected")
	}
}

func TestCreateTransactionWithBuildAccountIdInvalid(t *testing.T) {
	db, _ := repository.NewDBMock()
	accountRepository := repository.NewAccountRepository(db)
	accountSerivce := service.NewAccountService(accountRepository)
	transactionMock := repository.NewTransactionMock()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	usecase := NewTransactionUseCase(transactionService, accountSerivce)

	_, infraError, domainError := usecase.CreateTransaction("99c49b65-cc11-487f-864d-55dbb6c90a6", transactionMock.OperationType, transactionMock.Amount)

	if infraError != nil {
		t.Errorf("\nInfra error '%s' was not expected", infraError.Error())
	}

	if domainError == nil {
		t.Error("\nDomain error was expected")
	}
}
