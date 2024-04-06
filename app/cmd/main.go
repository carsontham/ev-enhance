package cmd

import "ev-enhance/app/cmd/api"

func main() {
	s := api.NewAPIServer(":3000")
	s.StartServer()
}
