package queryable

import (
	"database/sql"
	"log"
)

var DB *sql.DB

// OpenConnectionPool creates a package instance of a postgres connection pool
func OpenConnectionPool(dbConnString string) {
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
		panic(err)
	}

	DB = db

}
