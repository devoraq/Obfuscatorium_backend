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
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", d.UserHandler.RegisterUser)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", d.UserHandler.GetUser)       // GET /api/users/123
				r.Patch("/", d.UserHandler.UpdateUser)  // PATCH /api/users/123
				r.Delete("/", d.UserHandler.DeleteUser) //DELETE /api/users/123
			})
		})
	})

	return r
}
