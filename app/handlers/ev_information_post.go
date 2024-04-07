package handlers

import (
	"encoding/json"
	"ev-enhance/app/domain/model"
	"ev-enhance/app/infrastructure/repository"
	"net/http"
)

// PostEVInformation ...
func PostEVInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// Decode the JSON request body into RequestBody
		var requestBody model.EVTransaction

		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		repo := repository.NewEVTransactionRepository()
		response, err := repo.PostData(requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
}
