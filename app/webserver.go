package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/csazevedo/go-account-transaction/app/drive/database"
	"github.com/csazevedo/go-account-transaction/app/driven/webapi/action"
	"github.com/csazevedo/go-account-transaction/app/driven/webapi/middleware"
	"github.com/csazevedo/go-account-transaction/config"

	"github.com/gorilla/mux"
)

type webserver struct {
	router    *mux.Router
	dbAdapter *sql.DB
}

type handleRequestFunction func(dbAdapter *sql.DB, w http.ResponseWriter, r *http.Request)

func New(c *config.Config) *webserver {
	router := mux.NewRouter()
	adapeter := database.NewMySqlConnection(c.DBConfig)
	ws := &webserver{router: router, dbAdapter: adapeter}
	ws.routing()

	return ws
}

func (ws *webserver) routing() {
	v1 := ws.router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/health", ws.handleRequest(action.Health)).Methods(http.MethodGet)

	transactionRouter := v1.PathPrefix("/transactions").Subrouter()
	transactionRouter.HandleFunc("", ws.handleRequest(action.CreateTransaction)).Methods(http.MethodPost)
	transactionRouter.Use(
		middleware.NewLogger().Middleware,
		middleware.NewAuthorization().Middleware,
	)

	accountRouter := v1.PathPrefix("/accounts").Subrouter()
	accountRouter.HandleFunc("", ws.handleRequest(action.CreateAccount)).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{id:[0-9]+}", ws.handleRequest(action.GetAccount)).Methods(http.MethodGet)
	accountRouter.Use(
		middleware.NewLogger().Middleware,
		middleware.NewAuthorization().Middleware,
	)
}

func (ws *webserver) handleRequest(action handleRequestFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action(ws.dbAdapter, w, r)
	}
}

func (ws *webserver) Run(host string) {
	var wait time.Duration

	srv := &http.Server{
		Addr: host,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      ws.router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
