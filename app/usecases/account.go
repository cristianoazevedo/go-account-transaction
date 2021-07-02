package usecases

import (
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
)

type accountUseCase struct {
	service service.AccountService
}

//AccountUseCase interface representing the accountUseCase struct
type AccountUseCase interface {
	CreateAccount(documentNumber string) (model.Account, model.InfraError, model.DomainError)
	FindAccount(id string) (model.Account, model.InfraError, model.DomainError)
}

//NewAccountUseCase creates a new struct of account use case
func NewAccountUseCase(service service.AccountService) AccountUseCase {
	return &accountUseCase{
		service: service,
	}
}

//CreateAccount create a new account using the document number
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

//FindAccount search for an account using the account identifier
func (useCase *accountUseCase) FindAccount(id string) (model.Account, model.InfraError, model.DomainError) {
	accountID, err := model.BuildID(id)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	account, err := useCase.service.FindAccountByID(accountID)

	if err != nil {
		return nil, model.NewInfraError(err.Error()), nil
	}

	return account, nil, nil
}
