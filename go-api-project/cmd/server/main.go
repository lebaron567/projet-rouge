package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "go-api-project/internal/routes"
)

func main() {
    r := chi.NewRouter()

    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Setup routes
    routes.SetupRoutes()

    // Start the server
    port := ":8080"
    log.Printf("Starting server on %s", port)
    if err := http.ListenAndServe(port, r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}