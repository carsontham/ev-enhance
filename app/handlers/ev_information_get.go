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

		company, _ := repo.GetEVOperatorData2(operator)

		client := service.NewClient("https://83b53cac-faa4-434a-a366-f491c3c74384.mock.pstmn.io//ev")
		obj, _ := client.GetEVInformation(context.Background(), company.Operator_id, company.Charger_id)

		responseBody, err := json.Marshal(obj)
		if err != nil {
			return
		}

		w.Write(responseBody)

	}
}
