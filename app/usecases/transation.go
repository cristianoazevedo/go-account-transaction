package usecases

import (
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
)

type transactionUseCase struct {
	transactionService service.TransactionService
	accountSerivce     service.AccountService
}

//TransactonUseCase interface representing the transactionUseCase struct
type TransactonUseCase interface {
	CreateTransaction(accountID string, operationType int, amount float64) (model.Transaction, model.InfraError, model.DomainError)
}

//NewTransactionUseCase creates a new struct of transaction use case
func NewTransactionUseCase(transactionService service.TransactionService, accountSerivce service.AccountService) TransactonUseCase {
	return &transactionUseCase{
		transactionService: transactionService,
		accountSerivce:     accountSerivce,
	}
}

//CreateTransaction creates a transaction
func (useCase *transactionUseCase) CreateTransaction(accountID string, operationType int, amount float64) (model.Transaction, model.InfraError, model.DomainError) {
	id, err := model.BuildID(accountID)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	account, err := useCase.accountSerivce.FindAccountByID(id)

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

	transaction := model.NewTransaction(account, operationTypeModel, amountModel)

	if err := useCase.transactionService.CreateTransaction(transaction); err != nil {
		return nil, model.NewInfraError(err.Error()), nil
	}

	return transaction, nil, nil
}
