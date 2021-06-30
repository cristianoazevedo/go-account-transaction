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

func (usecase *createAcUsecase) Handle(document model.Document) (model.Account, model.DomainError, model.InfraError) {
	account, err := usecase.service.FindAccountByDocumentNumber(document)

	if err != nil {
		return nil, nil, model.NewInfraError(err.Error())
	}

	if account != nil {
		return nil, model.NewDomainError("account already exists"), nil
	}

	newAccount := model.NewAccount(document)

	createAccountError := usecase.service.CreateAccount(newAccount)

	if createAccountError != nil {
		return nil, nil, model.NewInfraError(createAccountError.Error())
	}

	return newAccount, nil, nil
}
