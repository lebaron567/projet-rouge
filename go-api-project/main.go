package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"api/config"
	"api/pkg/authentication"
	"api/pkg/comment"
	"api/pkg/like"
	"api/pkg/post"
	"api/pkg/user"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la configuration : %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // 🔒 à restreindre en prod (ex: "http://localhost:3000")
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "User-ID", "Current-User-ID"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", authentication.Routes(cfg))
		r.Mount("/comment", comment.Routes(cfg))
		r.Mount("/users", user.Routes(cfg))
		r.Mount("/like", like.Routes(cfg))
		r.Mount("/posts", post.Routes(cfg))
		r.Get("/users", user.GetAllUsers(cfg))
	})

	r.Group(func(r chi.Router) {
		r.Use(authentication.AuthMiddleware("c8f9d72e3b4a6d9e7f0b1c2a3e4f5g6h7i8j9k0l1m2n3o4p5q6r7s8t9u0v1w2x3"))

		r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
			user := authentication.GetUserFromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("Welcome, %s!", user)))
		})
	})

	// Servir la documentation Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Printf("Serveur démarré sur le port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
