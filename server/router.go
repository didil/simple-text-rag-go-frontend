package server

import "github.com/go-chi/chi/v5"

func NewRouter(app *App) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/collections", app.HandleCreateCollection)
		r.Post("/questions", app.HandleGetAnswer)
	})

	return r
}
