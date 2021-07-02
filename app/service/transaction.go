package service

import (
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
)

type tranactionService struct {
	repository repository.TransactionRepository
}

//TransactionService interface representing the transaction service struct
type TransactionService interface {
	CreateTransaction(transaction model.Transaction) error
}

//NewTransactionService creates a new struct of transaction service
func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &tranactionService{
		repository: repository,
	}
}

//CreateTransaction creates a transaction
func (tranactionService *tranactionService) CreateTransaction(transaction model.Transaction) error {
	return tranactionService.repository.CreateTransaction(transaction)
}
