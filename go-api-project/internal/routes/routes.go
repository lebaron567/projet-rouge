package routes

import (
    "github.com/go-chi/chi/v5"
    "go-api-project/internal/handlers"
)

func SetupRoutes() *chi.Mux {
    r := chi.NewRouter()

    r.Get("/resource", handlers.HandleGet)
    r.Post("/resource", handlers.HandlePost)

    return r
}