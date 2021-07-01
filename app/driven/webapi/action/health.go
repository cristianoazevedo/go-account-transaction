package action

import (
	"net/http"
)

type healthAction struct{}

func NewHealthAction() *healthAction {
	return &healthAction{}
}

func (action *healthAction) Health(w http.ResponseWriter, r *http.Request) {
	responder := NewResponder(w)
	response := ResponseInfo{Info: "alive"}

	responder.Accepted(response)
}
