package middleware

import (
	"net/http"

	"github.com/csazevedo/go-account-transaction/app/driven/webapi/action"
)

type authorization struct {
	auth string
}

func NewAuthorization() *authorization {
	return &authorization{auth: "0c7ee5a41bff7c8af4d4ff3740b0224d"}
}

func (authorization *authorization) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorization.auth == authorizationHeader {
			next.ServeHTTP(w, r)
		} else {
			responder := action.NewResponder(w)
			reponseError := action.ResponseError{Error: "authorization invalid"}
			responder.Forbidden(reponseError)
		}
	})
}
