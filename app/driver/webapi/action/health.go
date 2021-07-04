package action

import (
	"net/http"
)

type healthAction struct{}

//HealthAction interface representing the healthAction
type HealthAction interface {
	Health(w http.ResponseWriter, r *http.Request)
}

//NewHealthAction creates a new struct of health action
func NewHealthAction() HealthAction {
	return &healthAction{}
}

//Health public action for application health
func (action *healthAction) Health(w http.ResponseWriter, r *http.Request) {
	responder := NewResponder(w)
	response := ResponseInfo{Info: "alive"}

	responder.Accepted(response)
}
