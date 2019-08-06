package config

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_extra_golang")
		if err != nil {
			log.Fatal(err)
		}
	return db
}