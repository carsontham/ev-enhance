package api

import (
	"ev-enhance/app/handlers"

	"github.com/go-chi/chi/v5"
)

func (s *APIServer) SetupRoutes() {
	// Endpoints				Method 		Function		Description

	s.router.Route("/company", func(r chi.Router) {
		r.Post("/ev-information", handlers.GetEVInformation())
		r.Post("/ev-information", handlers.PostEVInformation())
	})
}
