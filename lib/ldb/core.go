package ldb

import (
	"database/sql"
	"log"

	"os"

	"github.com/gchaincl/dotsql"
	"github.com/joho/godotenv"

	// postgres lib used in connection
	_ "github.com/lib/pq"
)

type connection struct {
	dsn string
}

type dotMapper struct {
	resource *dotsql.DotSql
}

var mapper = &dotMapper{}
var objConn connection

// Connect to db
func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	objConn.dsn = os.Getenv("DB_DSN")

	db, err := sql.Open("postgres", objConn.dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// GenResource todo doc
func (d *dotMapper) GenResource(qpath string) {
	var dot, err = dotsql.LoadFromFile(qpath)

	if err != nil {
		log.Fatal(err)
	}

	d.resource = dot
}
