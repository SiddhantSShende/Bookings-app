package main

import (
	"net/http"

	"github.com/SiddhantSShende/bookings-app/pkg/config"
	"github.com/SiddhantSShende/bookings-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// Serving static files
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

/*
// routes sets up the application routes
func routes(app *config.AppConfig) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	return mux
}

*/
