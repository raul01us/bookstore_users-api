package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
	DB_HOST = "DB_HOST"
	DB_NAME = "DB_NAME"
)

var (
	Client *sql.DB

	dbUser = os.Getenv(DB_USER)
	dbPass = os.Getenv(DB_PASS)
	dbHost = os.Getenv(DB_HOST)
	dbName = os.Getenv(DB_NAME)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	// Create the connection to the server
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
