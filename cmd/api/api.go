package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config *config
}

type config struct {
	addr string
}

func (app *application) mount(mux *chi.Mux) error {
	mux.Use(middleware.Logger)
	mux.Route("/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("v1 OK"))
		})
	})
	return nil
}
