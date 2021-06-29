package service

import (
	"github.com/csazevedo/go-account-transaction/app/drive/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
)

type accountService struct {
	repository repository.AcccountRepository
}

type AccountService interface {
	CreateAccount(model.Account) error
}

func NewAccountService(repository repository.AcccountRepository) *accountService {
	return &accountService{
		repository: repository,
	}
}

func (service *accountService) CreateAccount(account model.Account) error {
	err := service.repository.CreateAccount(account)

	if err != nil {
		return err
	}

	return nil
}
