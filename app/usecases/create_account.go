package usecases

import (
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
)

type createAcUsecase struct {
	service service.AccountService
}

func NewCreateAccountUseCase(service service.AccountService) *createAcUsecase {
	return &createAcUsecase{
		service: service,
	}
}

func (usecase *createAcUsecase) Handle(document model.Document) (model.Account, error) {
	account, err := usecase.service.FindAccountByDocumentNumber(document)

	if err != nil {
		return nil, err
	}

	if account != nil {
		return nil, model.NewDomainError("account already exists")
	}

	newAccount := model.NewAccount(document)

	createAccountError := usecase.service.CreateAccount(newAccount)

	if createAccountError != nil {
		return nil, createAccountError
	}

	return newAccount, nil
}
