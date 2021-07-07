package usecases

import (
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
)

type transactionUseCase struct {
	transactionService service.TransactionService
	accountService     service.AccountService
}

//TransactonUseCase interface representing the transactionUseCase struct
type TransactonUseCase interface {
	CreateTransaction(accountID string, operationType int, amount float64) (model.Transaction, model.InfraError, model.DomainError)
}

//NewTransactionUseCase creates a new struct of transaction use case
func NewTransactionUseCase(transactionService service.TransactionService, accountService service.AccountService) TransactonUseCase {
	return &transactionUseCase{
		transactionService: transactionService,
		accountService:     accountService,
	}
}

//CreateTransaction creates a transaction
func (useCase *transactionUseCase) CreateTransaction(accountID string, operationType int, amount float64) (model.Transaction, model.InfraError, model.DomainError) {
	id, err := model.BuildID(accountID)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	account, err := useCase.accountService.FindAccountByID(id)

	if err != nil {
		return nil, model.NewInfraError(err.Error()), nil
	}

	if account == nil {
		return nil, nil, model.NewDomainError("account not found")
	}

	amountModel, err := model.NewAmount(amount)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	operationTypeModel, err := model.NewOperationType(operationType)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	transaction, err := model.NewTransaction(account, operationTypeModel, amountModel)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	if err := useCase.transactionService.CreateTransaction(transaction); err != nil {
		return nil, model.NewInfraError(err.Error()), nil
	}

	return transaction, nil, nil
}
