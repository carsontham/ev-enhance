package handlers

import (
	"context"
	"encoding/json"
	"ev-enhance/app/domain/model"
	"ev-enhance/app/infrastructure/repository"
	"ev-enhance/app/infrastructure/service"
	"fmt"
	"net/http"
)

// http://localhost:3000/company/ev-information
func GetEVInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var operator model.Operator
		if err := json.NewDecoder(req.Body).Decode(&operator); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("Get request is working..")
		repo := repository.NewEVTransactionRepository()
		//var result int
		//		result = repo.GetEVOperatorData(operator.operatorID)

		company, _ := repo.GetEVOperatorData2(operator)
		fmt.Println("company is ", company)
		client := service.NewClient("https://83b53cac-faa4-434a-a366-f491c3c74384.mock.pstmn.io//ev")
		obj, _ := client.GetEVInformation(context.Background(), company.Operator_id)
		// Serialize the dummy response to JSON
		responseBody, err := json.Marshal(obj)
		fmt.Println(responseBody)
		if err != nil {
			fmt.Println("Error marshalling dummy response:", err)
			return
		}

		w.Write(responseBody)

	}
}

// func sendResultToPostman(result int) {
// 	// Create a new client to interact with the mock Postman server
// 	postmanEndpoint := "https://83b53cac-faa4-434a-a366-f491c3c74384.mock.pstmn.io//ev"
// 	client := service.NewClient(postmanEndpoint)

// 	// Send the result to the mock Postman server
// 	err := client.SendIntegerToPostmanServer(result)
// 	if err != nil {
// 		fmt.Println("Error sending result to Postman server:", err)
// 	} else {
// 		fmt.Println("Result sent successfully to Postman server.")
// 	}
// }
