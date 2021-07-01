package action

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/csazevedo/go-account-transaction/app/drive/database/repository"
	"github.com/csazevedo/go-account-transaction/app/service"
	"github.com/csazevedo/go-account-transaction/app/usecases"
	"github.com/google/logger"
)

type transactionAction struct {
	dbAdapter  *sql.DB
	logAdapter *logger.Logger
}

type createTransactionBoby struct {
	AccountID     string  `json:"account_id"`
	OperationType int     `json:"operation_type"`
	Amount        float64 `json:"amount"`
}

type ResponseTransaction struct {
	TransactionID string `json:"transaction_id"`
}

func NewTransactionAction(dbAdapter *sql.DB, logAdapter *logger.Logger) *transactionAction {
	return &transactionAction{dbAdapter: dbAdapter, logAdapter: logAdapter}
}

func (action *transactionAction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var body createTransactionBoby
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

	tranaction, infraError, domainError := useCase.CreateTransaction(body.AccountID, body.OperationType, body.Amount)

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

	response := ResponseTransaction{TransactionID: tranaction.GetId().GetValue()}

	responder.Created(response)
}
