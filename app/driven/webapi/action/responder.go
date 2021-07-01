package action

import (
	"encoding/json"
	"net/http"
)

type responder struct {
	rw http.ResponseWriter
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseInfo struct {
	Info string `json:"info"`
}

func NewResponder(rw http.ResponseWriter) *responder {
	return &responder{rw: rw}
}

func (responder *responder) withJSON(status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		responder.rw.WriteHeader(http.StatusInternalServerError)
		responder.rw.Write([]byte(err.Error()))
		return
	}
	responder.rw.Header().Set("Content-Type", "application/json")
	responder.rw.WriteHeader(status)
	responder.rw.Write([]byte(response))
}

func (responder *responder) Accepted(payload interface{}) {
	responder.withJSON(http.StatusAccepted, payload)
}

func (responder *responder) OK(payload interface{}) {
	responder.withJSON(http.StatusOK, payload)
}

func (responder *responder) Created(payload interface{}) {
	responder.withJSON(http.StatusCreated, payload)
}

func (responder *responder) InternalServerError(payload interface{}) {
	responder.withJSON(http.StatusInternalServerError, payload)
}

func (responder *responder) BadRequest(payload interface{}) {
	responder.withJSON(http.StatusBadRequest, payload)
}

func (responder *responder) NotFound(payload interface{}) {
	responder.withJSON(http.StatusNotFound, payload)
}

func (responder *responder) Forbidden(payload interface{}) {
	responder.withJSON(http.StatusForbidden, payload)
}
