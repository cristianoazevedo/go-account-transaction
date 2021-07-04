package repository

import (
	"database/sql"

	"github.com/csazevedo/go-account-transaction/app/model"
)

type accountRepository struct {
	dbAdapter *sql.DB
}

//AcccountRepository interface representing the acccountRepository
type AcccountRepository interface {
	CreateAccount(account model.Account) error
	FindAccountByDocumentNumber(document model.Document) (model.Account, error)
	FindByID(idAccount model.ID) (model.Account, error)
}

//NewAccountRepository creates a new struct of acccount repository
func NewAccountRepository(adapter *sql.DB) AcccountRepository {
	return &accountRepository{dbAdapter: adapter}
}

//FindAccountByDocumentNumber find account by document number
func (repository *accountRepository) FindAccountByDocumentNumber(document model.Document) (model.Account, error) {
	query := "SELECT id, document_number, created_at FROM accounts where document_number = ?"

	return repository.find(query, document.GetValue())
}

//FindAccountByDocumentNumber find account by account ID
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

//CreateAccount create an account
//As a critical point, the concept of transaction is used.
//If there is any problem, no changes are made
func (repository *accountRepository) CreateAccount(account model.Account) (err error) {
	tx, err := repository.dbAdapter.Begin()

	if err != nil {
		return
	}

	_, err = tx.Exec("INSERT INTO accounts(id, document_number) VALUES(?, ?)", account.GetID().GetValue(), account.GetDocument().GetValue())

	if err != nil {
		tx.Rollback()

		return
	}

	err = tx.Commit()

	return
}
