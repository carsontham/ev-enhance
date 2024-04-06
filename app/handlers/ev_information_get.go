package handlers

import (
	"fmt"
	"net/http"
)

func GetEVInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Get request is working..")

	}
}
