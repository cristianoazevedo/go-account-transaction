package action

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/service"
	"github.com/csazevedo/go-account-transaction/app/usecases"
	"github.com/google/logger"
	"github.com/gorilla/mux"
)

type accountAction struct {
	dbAdapter  *sql.DB
	logAdapter *logger.Logger
}

//AccountAction interface representing the accountAction
type AccountAction interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
}

//NewAccountAction creates a new struct of acccount action
func NewAccountAction(dbAdapter *sql.DB, logAdapter *logger.Logger) AccountAction {
	return &accountAction{dbAdapter: dbAdapter, logAdapter: logAdapter}
}

type createAccountBody struct {
	DocumentNumber string `json:"document_number"`
}

type responseAccount struct {
	AccountID      string `json:"account_id"`
	DocumentNumber string `json:"document_number,omitempty"`
}

//CreateAccount action responsible for receiving a request and creating an account
func (action *accountAction) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var body createAccountBody
	responder := NewResponder(w)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		action.logAdapter.Errorf("Error to parse body: %s", err.Error())
		response := ResponseError{Error: err.Error()}
		responder.BadRequest(response)
		return
	}

	repository := repository.NewAccountRepository(action.dbAdapter)
	service := service.NewAccountService(repository)

	useCase := usecases.NewAccountUseCase(service)

	account, infraError, domainError := useCase.CreateAccount(body.DocumentNumber)

	if domainError != nil {
		action.logAdapter.Errorf("Error to create account: %s", domainError.Error())
		response := ResponseInfo{Info: domainError.Error()}
		responder.BadRequest(response)
		return
	}

	if infraError != nil {
		action.logAdapter.Errorf("Error to create account: %s", infraError.Error())
		response := ResponseError{Error: infraError.Error()}
		responder.InternalServerError(response)
		return
	}

	response := responseAccount{AccountID: account.GetID().GetValue()}

	responder.Created(response)
}

//CreateAccount action responsible for receiving a request and find an account
func (action *accountAction) GetAccount(w http.ResponseWriter, r *http.Request) {
	id, found := mux.Vars(r)["id"]

	responder := NewResponder(w)

	if !found {
		action.logAdapter.Error("Parameter id not informed")
		response := ResponseError{Error: "parameter id not informed"}
		responder.BadRequest(response)
		return
	}

	repository := repository.NewAccountRepository(action.dbAdapter)
	service := service.NewAccountService(repository)

	useCase := usecases.NewAccountUseCase(service)

	account, infraError, domainError := useCase.FindAccount(id)

	if domainError != nil {
		action.logAdapter.Errorf("Error to find account: %s", domainError.Error())
		response := ResponseInfo{Info: domainError.Error()}
		responder.BadRequest(response)
		return
	}

	if infraError != nil {
		action.logAdapter.Errorf("Error to find account: %s", infraError.Error())
		response := ResponseError{Error: infraError.Error()}
		responder.InternalServerError(response)
		return
	}

	if account == nil {
		reponse := ResponseInfo{Info: "account not found"}
		responder.NotFound(reponse)
		return
	}

	response := responseAccount{AccountID: account.GetID().GetValue(), DocumentNumber: account.GetDocument().GetValue()}

	responder.OK(response)
}
