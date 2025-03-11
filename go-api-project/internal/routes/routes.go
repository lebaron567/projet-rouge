package routes

import (
    "github.com/go-chi/chi/v5"
    "go-api-project/internal/handlers"
)

func SetupRoutes() *chi.Mux {
    r := chi.NewRouter()

    r.Get("/", handlers.HandleGet)
    r.Post("/", handlers.HandlePost)

    return r
}