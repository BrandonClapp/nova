package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// OpenConnectionPool creates a package instance of a postgres connection pool
func OpenConnectionPool(dbConnString string) {
	gormDb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = gormDb
}
