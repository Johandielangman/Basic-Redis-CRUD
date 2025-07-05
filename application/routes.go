package application

import (
	"net/http"

	"github.com/Johandielangman/Basic-Redis-CRUD/handler"
	"github.com/Johandielangman/Basic-Redis-CRUD/repository/order"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}

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
