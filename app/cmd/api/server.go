package api

import (
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
	http.ListenAndServe(s.listenAddr, s.router)
}
