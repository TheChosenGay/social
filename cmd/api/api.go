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
	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	return nil
}
