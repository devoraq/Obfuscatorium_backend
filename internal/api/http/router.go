package api

import (
	"net/http"

	"github.com/devoraq/Obfuscatorium_backend/internal/api/http/handlers"
	"github.com/go-chi/chi/v5"
)

type Deps struct {
	UserHandler *handlers.UserHandler
}

func NewRouter(d Deps) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Post("/register", d.UserHandler.RegisterUser)
			r.Get("/get", d.UserHandler.GetUser)
		})
	})

	return r
}
