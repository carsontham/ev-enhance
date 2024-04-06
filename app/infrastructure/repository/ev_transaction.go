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

func (r EVTransactionRepository) GetEVOperatorData(charger_id string) int {
	// Call the getDB function to obtain the database connection
	db := r.getDB()

	// Check for errors
	if db == nil {
		log.Fatal("Database connection is nil")
	}

	row, err := db.Query("SELECT operator_id FROM chargers WHERE charger_id = ?", charger_id)
	if err != nil {
		log.Fatal("Error while querying the db: ", err)
	}
	defer row.Close()

	var operator_id int


	// Scan the values from the row into variables
	if err := row.Scan(&operator_id); err != nil {
		log.Fatal("Error scanning row: ", err)
	}
		// Print the values
	fmt.Printf("Operator ID: %d", operator_id)
	fmt.Println("Finished Operation from GetData()")

	log.Println("Successfully retrieved data from the database")
	return operator_id
}
