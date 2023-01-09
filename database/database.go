package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

// Connect opens a connection to the database and stores the database handle in the database variable.
func Connect() {
	var err error
	database, err = sql.Open("sqlite3", "/var/imunify360/imunify360.db")
	if err != nil {
		panic(err)
	}
}
