package main

import (
	"go-expert-otel/http-server/infra/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handlers.IndexHandler)
	router.Post("/temperature", handlers.TemperatureHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
