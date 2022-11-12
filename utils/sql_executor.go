package utils

import (
	"database/sql"
	"log"
)

func ExecuteSql(db *sql.DB, sql string) {
	result, err := db.Exec(string(sql))
	if err != nil {
		log.Printf("error while executing sql: %v", err)
		panic(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Print("error reading sql rows affected")
	}

	// log.Printf("sql ran successful. %v rows affected", affected)
}
