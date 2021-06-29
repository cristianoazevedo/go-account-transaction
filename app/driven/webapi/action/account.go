package action

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/csazevedo/go-account-transaction/app/drive/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
	"github.com/csazevedo/go-account-transaction/app/usecases"
)

type CreateAccountBody struct {
	DocumentNumber string `json:"document_number"`
}

type ResponseCreateAccount struct {
	AccountID string `json:"account_id"`
}

func CreateAccount(dbAdapter *sql.DB, w http.ResponseWriter, r *http.Request) {
	var body CreateAccountBody
	responder := NewResponder(w)
	repository := repository.NewAccountRepository(dbAdapter)
	service := service.NewAccountService(repository)

	useCase := usecases.NewCreateAccountUseCase(service)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		reponseError := ResponseError{Error: err.Error()}
		responder.badRequest(reponseError)
		return
	}

	document := model.NewDocument(body.DocumentNumber)

	account, err := useCase.Handle(document)

	if err != nil {
		reponseError := ResponseError{Error: err.Error()}
		switch err.(type) {
		case model.DomainError:
			responder.badRequest(reponseError)
		default:
			responder.internalServerError(reponseError)
		}

		return
	}

	responseCreateAccount := ResponseCreateAccount{AccountID: account.GetId().GetValue()}

	responder.created(responseCreateAccount)
}

func GetAccount(dbAdapter *sql.DB, w http.ResponseWriter, r *http.Request) {
	responder := NewResponder(w)
	health := map[string]string{}
	health["foo"] = "bar"
	health["xoo"] = "berr"
	responder.accepted(health)
}
