package repository

import (
	"database/sql"
	"ev-enhance/app/domain/model"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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

func (r EVTransactionRepository) PostData(arg model.EVTransaction) (model.EVTransaction, error) {

	const createTransaction = `
	INSERT INTO transactions (
		user_id,
		operator_id,
		point_id,
		charging_time,
		payment_type,
		status
	) 
	VALUES (
		$1, $2, $3, $4, $5, $6
	)
	RETURNING id, user_id, operator_id, point_id, charging_time, payment_type, status
	`

	// Call the getDB function to obtain the database connection
	db := r.getDB()
	// Check for errors
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	row := db.QueryRow(createTransaction,
		arg.User_id,
		arg.Operator_id,
		arg.Point_id,
		arg.Charging_time,
		arg.Payment_type,
		arg.Status,
	)
	var i model.EVTransaction
	err := row.Scan(
		&i.Id,
		&i.User_id,
		&i.Operator_id,
		&i.Point_id,
		&i.Charging_time,
		&i.Payment_type,
		&i.Status,
	)
	return i, err
}
