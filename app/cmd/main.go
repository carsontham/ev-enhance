package main

import (
	"ev-enhance/app/cmd/api"
	"ev-enhance/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Some error occured. Err: ", err)
	}

	db.InitDB()
	defer db.CloseDB()

	s := api.NewAPIServer(":3000")
	s.SetupRoutes()
	s.StartServer()
}
