package action

import (
	"database/sql"
	"net/http"
)

type transactionAction struct {
	dbAdapter *sql.DB
}

func NewTransactionAction(dbAdapter *sql.DB) *transactionAction {
	return &transactionAction{dbAdapter: dbAdapter}
}

func (action *transactionAction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	responder := NewResponder(w)
	health := map[string]string{}
	health["foo"] = "bar"
	health["xoo"] = "berr"
	responder.created(health)
}
