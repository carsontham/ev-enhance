package handlers

import (
	"encoding/json"
	"ev-enhance/app/infrastructure/repository"
	"fmt"
	"net/http"
)

// http://localhost:3000/company/ev-information
func GetEVInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("Get request is working..")
		repo := repository.NewEVTransactionRepository()
		repo.GetData()

		// TODO: take in chargerID, find the operator from DB
		// In request, need to pass correct operator and chargerID
		// $0.80 / KWH -> get from mock server

		type DummyResponse struct {
			Message string `json:"message"`
			Data    []int  `json:"data"`
		}

		dummy := DummyResponse{
			Message: "Dummy response message",
			Data:    []int{1, 2, 3, 4, 5},
		}

		// Serialize the dummy response to JSON
		responseBody, err := json.Marshal(dummy)
		if err != nil {
			fmt.Println("Error marshalling dummy response:", err)
			return
		}

		w.Write(responseBody)

	}
}
