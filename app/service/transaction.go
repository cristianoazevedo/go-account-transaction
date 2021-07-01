package service

import (
	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
)

type tranactionService struct {
	repository repository.TransactionRepository
}

type TransactionService interface {
	CreateTransaction(transaction model.Transaction) error
}

func NewTransactionService(repository repository.TransactionRepository) *tranactionService {
	return &tranactionService{
		repository: repository,
	}
}

func (tranactionService *tranactionService) CreateTransaction(transaction model.Transaction) error {
	return tranactionService.repository.CreateTransaction(transaction)
}
