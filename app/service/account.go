package service

import (
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
)

type accountService struct {
	repository repository.AcccountRepository
}

//AccountService interface representing the account service struct
type AccountService interface {
	CreateAccount(account model.Account) error
	FindAccountByDocumentNumber(document model.Document) (model.Account, error)
	FindAccountByID(accountID model.ID) (model.Account, error)
}

//NewAccountService creates a new struct of acccount service
func NewAccountService(repository repository.AcccountRepository) AccountService {
	return &accountService{
		repository: repository,
	}
}

//FindAccountByDocumentNumber search for an account using the document number
func (service *accountService) FindAccountByDocumentNumber(document model.Document) (model.Account, error) {
	return service.repository.FindAccountByDocumentNumber(document)
}

//FindAccount search for an account using the account identifier
func (service *accountService) FindAccountByID(accountID model.ID) (model.Account, error) {
	return service.repository.FindByID(accountID)
}

//CreateAccount create a new account using the document number
func (service *accountService) CreateAccount(account model.Account) error {
	return service.repository.CreateAccount(account)
}
