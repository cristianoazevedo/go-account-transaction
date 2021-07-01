package service

import (
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
)

type accountService struct {
	repository repository.AcccountRepository
}

type AccountService interface {
	CreateAccount(account model.Account) error
	FindAccountByDocumentNumber(document model.Document) (model.Account, error)
	FindAccountByID(accountID model.ID) (model.Account, error)
}

func NewAccountService(repository repository.AcccountRepository) *accountService {
	return &accountService{
		repository: repository,
	}
}

func (service *accountService) FindAccountByDocumentNumber(document model.Document) (model.Account, error) {
	return service.repository.FindAccountByDocumentNumber(document)
}

func (service *accountService) FindAccountByID(accountID model.ID) (model.Account, error) {
	return service.repository.FindByID(accountID)
}

func (service *accountService) CreateAccount(account model.Account) error {
	return service.repository.CreateAccount(account)
}
