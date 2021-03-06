package middleware

import (
	"net/http"

	"github.com/google/logger"
)

type loggerRequest struct {
	logAdapter *logger.Logger
}

//LoggerRequest interface representing the loggerRequest struct
type LoggerRequest interface {
	Middleware(next http.Handler) http.Handler
}

//NewLoggerRequest creates a new struct of loggerRequest middleware
func NewLoggerRequest(logAdapter *logger.Logger) LoggerRequest {
	return &loggerRequest{logAdapter: logAdapter}
}

//Middleware loggerRequest middleware handler
func (lr *loggerRequest) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lr.logAdapter.Infof("Request URI %s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
