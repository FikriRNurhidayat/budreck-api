package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(connectionURL string) (err error) {
	db, err = sql.Open("postgres", connectionURL)
	db.Ping()
	return
}
