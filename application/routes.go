package application

import (
	"net/http"

	"chi2/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	// ====> C - CREATE
	router.Post("/", orderHandler.Create)

	// ====> R - READ
	router.Get("/{id}", orderHandler.GetByID)
	router.Get("/", orderHandler.List)

	// ====> U - UPDATE
	router.Put("/{id}", orderHandler.UpdateByID)

	// ====> D - DELETE
	router.Delete("/{id}", orderHandler.DeleteByID)
}
