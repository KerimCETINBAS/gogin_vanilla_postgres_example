package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "tutorial"
)

var psqlconn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func InitDatabase() {
	if db, err := sql.Open("postgres", psqlconn); err != nil {
		panic("Can not connect to database")
	} else {
		DB = db
	}
}
