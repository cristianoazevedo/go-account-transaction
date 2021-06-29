package action

import (
	"database/sql"
	"net/http"
)

func Health(dbAdapter *sql.DB, w http.ResponseWriter, r *http.Request) {
	responder := NewResponder(w)
	health := map[string]bool{}
	health["alive"] = true
	responder.accepted(health)
}
