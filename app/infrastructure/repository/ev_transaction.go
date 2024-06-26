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

func (r EVTransactionRepository) GetEVOperatorData2(arg model.Operator) (model.Operators, error) {
	const getCharger = `
	SELECT 
	charger_id,
	chargers.operator_id, 
	location,
	operators.name
FROM chargers JOIN operators ON chargers.operator_id = operators.operator_id
WHERE charger_id = $1
	`
	// Call the getDB function to obtain the database connection
	db := r.getDB()
	// Check for errors
	if db == nil {
		return model.Operators{}, fmt.Errorf("database connection is nil")
	}
	row := db.QueryRow(getCharger, arg.Charger_id)
	var i model.Operators
	err := row.Scan(
		&i.Charger_id,
		&i.Operator_id,
		&i.Location,
		&i.Name,
	)
	return i, err
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
