package comment

import (
	"api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/{id}/comments", AddComment(cfg))
	r.Get("/{id}/comments", GetComments(cfg))
	r.Delete("/{id}", DeleteComment(cfg))
	return r
}
