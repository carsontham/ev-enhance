package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type APIServer struct {
	listenAddr string
	router     chi.Router
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		router:     chi.NewRouter(),
	}
}

func (s *APIServer) StartServer() {
	fmt.Println("Starting server... listening on port :3000")
	http.ListenAndServe(s.listenAddr, s.router)
	fmt.Println("Starting server... listening on port :3000")

}
