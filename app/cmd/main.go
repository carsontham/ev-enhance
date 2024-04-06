package main

import (
	"ev-enhance/app/cmd/api"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Some error occured. Err: ", err)
	}

	s := api.NewAPIServer(":3000")
	s.SetupRoutes()
	s.StartServer()
}
