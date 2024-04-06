package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Some error occured. Err: ", err)
	}

	// setup db connection
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	defer db.Close()

	// Perform a sample query
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Error while querying the db: ", err)
	}
	defer rows.Close()
	fmt.Println(rows)

}
