package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config := &config{
		addr: ":8080",
	}
	app := &application{
		config: config,
	}
	mux := chi.NewRouter()
	app.mount(mux)
	server := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}
	server.ListenAndServe()
}
