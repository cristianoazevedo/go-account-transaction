package database

import (
	"database/sql"

	"github.com/csazevedo/go-account-transaction/config"

	"github.com/go-sql-driver/mysql"
)

//NewMySQLConnection create a data base connection
func NewMySQLConnection(config *config.DBConfig) *sql.DB {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = config.User
	mysqlConfig.Passwd = config.Pwd
	mysqlConfig.Addr = config.Host
	mysqlConfig.DBName = config.Name
	mysqlConfig.Net = "tcp"

	db, err := sql.Open(config.Driver, mysqlConfig.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	return db
}
