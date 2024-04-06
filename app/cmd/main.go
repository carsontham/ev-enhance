package main

import "ev-enhance/app/cmd/api"

func main() {
	s := api.NewAPIServer(":3000")
	s.SetupRoutes()
	s.StartServer()
}
