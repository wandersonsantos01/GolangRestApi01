package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	connStr := "user=root dbname=market password=root host=host.docker.internal sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
