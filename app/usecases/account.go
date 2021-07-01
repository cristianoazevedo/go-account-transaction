package usecases

import (
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
)

type accountUseCase struct {
	service service.AccountService
}

func NewAccountUseCase(service service.AccountService) *accountUseCase {
	return &accountUseCase{
		service: service,
	}
}

func (useCase *accountUseCase) CreateAccount(documentNumber string) (model.Account, model.InfraError, model.DomainError) {
	document, err := model.NewDocument(documentNumber)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	account, err := useCase.service.FindAccountByDocumentNumber(document)

	if err != nil {
		return nil, model.NewInfraError(err.Error()), nil
	}

	if account != nil {
		return nil, nil, model.NewDomainError("account already exists")
	}

	accountModel := model.NewAccount(document)

	if err := useCase.service.CreateAccount(accountModel); err != nil {
		return nil, model.NewInfraError(err.Error()), nil
	}

	return accountModel, nil, nil
}

func (useCase *accountUseCase) FindAccount(id string) (model.Account, model.InfraError, model.DomainError) {
	accountID, err := model.BuildId(id)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	account, err := useCase.service.FindAccountByID(accountID)

	if err != nil {
		return nil, model.NewInfraError(err.Error()), nil
	}

	return account, nil, nil
}
