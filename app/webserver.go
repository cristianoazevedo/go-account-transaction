package app

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/csazevedo/go-account-transaction/app/driven/database"
	logbook "github.com/csazevedo/go-account-transaction/app/driven/log"
	"github.com/csazevedo/go-account-transaction/app/driver/webapi/action"
	"github.com/csazevedo/go-account-transaction/app/driver/webapi/middleware"
	"github.com/csazevedo/go-account-transaction/config"
	"github.com/google/logger"

	"github.com/gorilla/mux"
)

type webserver struct {
	router     *mux.Router
	dbAdapter  *sql.DB
	logAdapter *logger.Logger
}

type handleRequestFunction func(dw http.ResponseWriter, r *http.Request)

//New bootstrap of application
//Defines application routes
//Create a database connection
//Set the log
func New(c *config.Config) *webserver {
	router := mux.NewRouter()
	dbAdapeter := database.NewMySQLConnection(c.DBConfig)
	logAdapter := logbook.NewLogger()
	ws := &webserver{router: router, dbAdapter: dbAdapeter, logAdapter: logAdapter}
	ws.routing()

	return ws
}

func (ws *webserver) routing() {
	healthAction := action.NewHealthAction()
	transactionAction := action.NewTransactionAction(ws.dbAdapter, ws.logAdapter)
	accountAction := action.NewAccountAction(ws.dbAdapter, ws.logAdapter)

	v1 := ws.router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/health", ws.handleRequest(healthAction.Health)).Methods(http.MethodGet)
	v1.Use(
		middleware.NewLoggerRequest(ws.logAdapter).Middleware,
	)

	transactionRouter := v1.PathPrefix("/transactions").Subrouter()
	transactionRouter.HandleFunc("", ws.handleRequest(transactionAction.CreateTransaction)).Methods(http.MethodPost)
	transactionRouter.Use(
		middleware.NewAuthorization(ws.logAdapter).Middleware,
	)

	accountRouter := v1.PathPrefix("/accounts").Subrouter()
	accountRouter.HandleFunc("", ws.handleRequest(accountAction.CreateAccount)).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{id:[a-z-0-9]+}", ws.handleRequest(accountAction.GetAccount)).Methods(http.MethodGet)
	accountRouter.Use(
		middleware.NewAuthorization(ws.logAdapter).Middleware,
	)
}

func (ws *webserver) handleRequest(handle handleRequestFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handle(w, r)
	}
}

//Run start webserver
func (ws *webserver) Run(host string) {
	srv := &http.Server{
		Addr:         host,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      ws.router,
	}

	if err := srv.ListenAndServe(); err != nil {
		ws.logAdapter.Fatalf("Failed start server: %v", err)
	}
}
