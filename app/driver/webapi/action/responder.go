package action

import (
	"encoding/json"
	"net/http"
)

type responder struct {
	rw http.ResponseWriter
}

//Responder interface representing the responder struct
type Responder interface {
	Accepted(payload interface{})
	OK(payload interface{})
	Created(payload interface{})
	InternalServerError(payload interface{})
	BadRequest(payload interface{})
	NotFound(payload interface{})
	Forbidden(payload interface{})
}

//ResponseError struct representing an error response
type ResponseError struct {
	Error string `json:"error"`
}

//ResponseInfo struct representing an information response
type ResponseInfo struct {
	Info string `json:"info"`
}

//NewResponder creates a new struct of responder
func NewResponder(rw http.ResponseWriter) Responder {
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

//Accepted return json with status code 202
func (responder *responder) Accepted(payload interface{}) {
	responder.withJSON(http.StatusAccepted, payload)
}

//OK return json with status code 200
func (responder *responder) OK(payload interface{}) {
	responder.withJSON(http.StatusOK, payload)
}

//Created return json with status code 201
func (responder *responder) Created(payload interface{}) {
	responder.withJSON(http.StatusCreated, payload)
}

//InternalServerError return json with status code 500
func (responder *responder) InternalServerError(payload interface{}) {
	responder.withJSON(http.StatusInternalServerError, payload)
}

//InternalServerError return json with status code 500
func (responder *responder) BadRequest(payload interface{}) {
	responder.withJSON(http.StatusBadRequest, payload)
}

//NotFound return json with status code 404
func (responder *responder) NotFound(payload interface{}) {
	responder.withJSON(http.StatusNotFound, payload)
}

//Forbidden return json with status code 403
func (responder *responder) Forbidden(payload interface{}) {
	responder.withJSON(http.StatusForbidden, payload)
}
