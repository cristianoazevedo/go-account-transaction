package usecases

import (
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
)

type transactionUseCase struct {
	transactionService service.TransactionService
	accountSerivce     service.AccountService
}

func NewTransactionUseCase(transactionService service.TransactionService, accountSerivce service.AccountService) *transactionUseCase {
	return &transactionUseCase{
		transactionService: transactionService,
		accountSerivce:     accountSerivce,
	}
}

func (useCase *transactionUseCase) CreateTransaction(accountID string, operationType int, amount float64) (model.Transaction, model.InfraError, model.DomainError) {
	id, err := model.BuildId(accountID)

	if err != nil {
		return nil, nil, model.NewDomainError(err.Error())
	}

	account, err := useCase.accountSerivce.FindAccountByID(id)

	if err != nil {
		return nil, model.NewInfraError(err.Error()), nil
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
