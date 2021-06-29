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
	account := model.NewAccount(document)

	err := usecase.service.CreateAccount(account)

	if err != nil {
		return nil, err
	}

	return account, nil
}
