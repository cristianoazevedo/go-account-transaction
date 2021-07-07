package action

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/csazevedo/go-account-transaction/app/driven/database/repository"
	"github.com/csazevedo/go-account-transaction/app/service"
	"github.com/csazevedo/go-account-transaction/app/usecases"
	"github.com/google/logger"
)

type transactionAction struct {
	dbAdapter  *sql.DB
	logAdapter *logger.Logger
}

//TransactionAction interface representing the transactionAction
type TransactionAction interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

type createTransactionBody struct {
	AccountID     string  `json:"account_id"`
	OperationType int     `json:"operation_type"`
	Amount        float64 `json:"amount"`
}

type responseTransaction struct {
	TransactionID string `json:"transaction_id"`
}

//NewTransactionAction creates a new struct of transaction action
func NewTransactionAction(dbAdapter *sql.DB, logAdapter *logger.Logger) TransactionAction {
	return &transactionAction{dbAdapter: dbAdapter, logAdapter: logAdapter}
}

//CreateAccount action responsible for receiving a request and creating an transaction
func (action *transactionAction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var body createTransactionBody
	responder := NewResponder(w)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		action.logAdapter.Errorf("Error to parse body: %s", err.Error())
		response := ResponseError{Error: err.Error()}
		responder.BadRequest(response)
		return
	}

	accountRepository := repository.NewAccountRepository(action.dbAdapter)
	accountService := service.NewAccountService(accountRepository)

	transactionRepository := repository.NewTransactionRepository(action.dbAdapter)
	transactionService := service.NewTransactionService(transactionRepository)

	useCase := usecases.NewTransactionUseCase(transactionService, accountService)

	transaction, infraError, domainError := useCase.CreateTransaction(body.AccountID, body.OperationType, body.Amount)

	if domainError != nil {
		action.logAdapter.Errorf("Error to create transaction: %s", domainError.Error())
		response := ResponseInfo{Info: domainError.Error()}
		responder.BadRequest(response)
		return
	}

	if infraError != nil {
		action.logAdapter.Errorf("Error to create transaction: %s", infraError.Error())
		response := ResponseError{Error: infraError.Error()}
		responder.InternalServerError(response)
		return
	}

	response := responseTransaction{TransactionID: transaction.GetID().GetValue()}

	responder.Created(response)
}
