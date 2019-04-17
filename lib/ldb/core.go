package ldb

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type connection struct {
	dsn string
}

var objConn connection

//Connect to db
func Connect() *sql.DB {
	objConn.dsn = "postgres://<USER>:<PASSWORD>@127.0.0.1/<DBNAME>"
	db, err := sql.Open("postgres", objConn.dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
