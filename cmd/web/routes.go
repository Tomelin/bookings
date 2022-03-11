package main

import (
	"net/http"

	"github.com/tomelin/bookings/pkg/config"
	"github.com/tomelin/bookings/pkg/handlers"

	"github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routesPat(app *config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
