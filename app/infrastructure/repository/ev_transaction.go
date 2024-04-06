package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type GetDB func() *sql.DB

type EVTransactionRepository struct {
	getDB GetDB
}

func NewEVTransactionRepository() *EVTransactionRepository {
	return &EVTransactionRepository{
		getDB: GetDBConnection,
	}
}

func GetDBConnection() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	return db
}

func (r EVTransactionRepository) GetData() {
	// Call the getDB function to obtain the database connection
	db := r.getDB()

	// Check for errors
	if db == nil {
		log.Fatal("Database connection is nil")
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Error while querying the db: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userID int
		var username string
		var email string
		var pw string
		var address string

		// Scan the values from the row into variables
		if err := rows.Scan(&userID, &username, &email, &pw, &address); err != nil {
			log.Fatal("Error scanning row: ", err)
		}
		// Print the values
		fmt.Printf("User ID: %d, Username: %s, Email: %s\n", userID, username, email)
		fmt.Println("out")
	}

	log.Println("Successfully retrieved data from the database")
}
