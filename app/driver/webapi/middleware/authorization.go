package middleware

import (
	"net/http"

	"github.com/csazevedo/go-account-transaction/app/driver/webapi/action"
	"github.com/google/logger"
)

type authorization struct {
	auth       string
	logAdapter *logger.Logger
}

//NewAuthorization creates a new struct of authorization middleware
func NewAuthorization(logAdapter *logger.Logger) *authorization {
	return &authorization{auth: "0c7ee5a41bff7c8af4d4ff3740b0224d", logAdapter: logAdapter}
}

//Middleware authorization middleware handler
func (authorization *authorization) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorization.auth == authorizationHeader {
			next.ServeHTTP(w, r)
			return
		}

		responder := action.NewResponder(w)

		if authorizationHeader == "" {
			response := action.ResponseInfo{Info: "authorization missing"}
			responder.BadRequest(response)
			return
		}

		authorization.logAdapter.Infof("Authorization invalid: %s", authorizationHeader)
		reponseError := action.ResponseError{Error: "authorization invalid"}
		responder.Forbidden(reponseError)
	})
}
