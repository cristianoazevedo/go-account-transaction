package repository

import (
	"database/sql"
	"database/sql/driver"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

//IDMock mock for test
type IDMock struct{}

// Match satisfies sqlmock.Argument interface
func (idMock IDMock) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

type accountMock struct {
	ID             string
	CreatedAt      string
	DocumentNumber string
}

//NewAccountMock mock for test
func NewAccountMock() *accountMock {
	return &accountMock{
		ID:             uuid.New().String(),
		CreatedAt:      "2021-01-01 00:00:00",
		DocumentNumber: "03393983024",
	}
}

type transactionMock struct {
	ID            string
	Amount        float64
	OperationType int
}

//NewTransactionMock mock for test
func NewTransactionMock() *transactionMock {
	return &transactionMock{
		ID:            uuid.New().String(),
		Amount:        10.0,
		OperationType: 1,
	}
}

//NewDBMock mock data base for test
func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
