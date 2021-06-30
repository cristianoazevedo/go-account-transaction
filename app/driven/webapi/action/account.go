package action

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/csazevedo/go-account-transaction/app/drive/database/repository"
	"github.com/csazevedo/go-account-transaction/app/model"
	"github.com/csazevedo/go-account-transaction/app/service"
	"github.com/csazevedo/go-account-transaction/app/usecases"
	"github.com/google/logger"
)

type accountAction struct {
	dbAdapter  *sql.DB
	logAdapter *logger.Logger
}

func NewAccountAction(dbAdapter *sql.DB, logAdapter *logger.Logger) *accountAction {
	return &accountAction{dbAdapter: dbAdapter, logAdapter: logAdapter}
}

type CreateAccountBody struct {
	DocumentNumber string `json:"document_number"`
}

type ResponseCreateAccount struct {
	AccountID string `json:"account_id"`
}

func (action *accountAction) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var body CreateAccountBody
	responder := NewResponder(w)
	repository := repository.NewAccountRepository(action.dbAdapter)
	service := service.NewAccountService(repository)

	useCase := usecases.NewCreateAccountUseCase(service)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		action.logAdapter.Errorf("error to parse body: %s", err.Error())
		reponseError := ResponseError{Error: err.Error()}
		responder.badRequest(reponseError)
		return
	}

	document := model.NewDocument(body.DocumentNumber)

	account, domainError, infraError := useCase.Handle(document)

	if domainError != nil {
		action.logAdapter.Errorf("Error to create account: %s", domainError.Error())
		reponseError := ResponseError{Error: domainError.Error()}
		responder.badRequest(reponseError)
		return
	}

	if infraError != nil {
		action.logAdapter.Errorf("Error to create account: %s", infraError.Error())
		reponseError := ResponseError{Error: infraError.Error()}
		responder.internalServerError(reponseError)
		return
	}

	responseCreateAccount := ResponseCreateAccount{AccountID: account.GetId().GetValue()}

	responder.created(responseCreateAccount)
}

func (action *accountAction) GetAccount(w http.ResponseWriter, r *http.Request) {
	responder := NewResponder(w)
	health := map[string]string{}
	health["foo"] = "bar"
	health["xoo"] = "berr"
	responder.accepted(health)
}
