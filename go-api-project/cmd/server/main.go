package main

import (
	"log"
	"net/http"

	"go-api-project/internal/db"
	"go-api-project/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db.Init()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Setup routes
	r = routes.SetupRoutes()

	// Start the server
	port := ":8088"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
