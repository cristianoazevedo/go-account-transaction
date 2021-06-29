package action

import (
	"database/sql"
	"net/http"
)

func CreateTransaction(dbAdapter *sql.DB, w http.ResponseWriter, r *http.Request) {
	responder := NewResponder(w)
	health := map[string]string{}
	health["foo"] = "bar"
	health["xoo"] = "berr"
	responder.created(health)
}
