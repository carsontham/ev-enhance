package db

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

// setup db connection
func InitDB() {
	db, err := sql.Open(os.Getenv("DBDriver"), os.Getenv("DBSource"))
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	defer db.Close()
}

// close db connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
