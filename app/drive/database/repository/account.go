package repository

import (
	"context"
	"database/sql"

	"github.com/csazevedo/go-account-transaction/app/model"
)

type accountRepository struct {
	dbAdapter *sql.DB
}

type AcccountRepository interface {
	CreateAccount(account model.Account) error
	FindAccountByDocumentNumber(document model.Document) (model.Account, error)
	FindByID(idAccount model.ID) (model.Account, error)
}

func NewAccountRepository(adapter *sql.DB) *accountRepository {
	return &accountRepository{dbAdapter: adapter}
}

func (repository *accountRepository) FindAccountByDocumentNumber(document model.Document) (model.Account, error) {
	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"

	return repository.find(query, document.GetValue())
}

func (repository *accountRepository) FindByID(idAccount model.ID) (model.Account, error) {
	query := "SELECT id, document_number, created_at FROM accounts where id = ?"

	return repository.find(query, idAccount.GetValue())
}

func (repository *accountRepository) find(queryString string, args ...interface{}) (model.Account, error) {
	query := repository.dbAdapter.QueryRow(queryString, args...)

	var id, documentNumber, createAt string

	err := query.Scan(&id, &documentNumber, &createAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	account, err := model.BuildAccount(id, documentNumber, createAt)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repository *accountRepository) CreateAccount(account model.Account) error {
	ctx := context.Background()

	transaction, err := repository.dbAdapter.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	insert, err := repository.dbAdapter.Prepare("INSERT INTO accounts(id, document_number) VALUES(?,?)")

	if err != nil {
		return err
	}

	_, err = insert.ExecContext(ctx, account.GetId().GetValue(), account.GetDocument().GetValue())

	if err != nil {
		transaction.Rollback()

		return err
	}

	err = transaction.Commit()

	if err != nil {
		return err
	}

	return nil
}
