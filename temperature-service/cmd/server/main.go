package main

import (
	"go-exper-otel/temperature/configs"
	"go-exper-otel/temperature/infra/webserver/handlers"
	"net/http"
)

func main() {
	config := configs.LoadConfig(".")

	temperatureHandler := handlers.NewTemperatureHandler(config)
	http.HandleFunc("/temperature", temperatureHandler.TemperatureHandler)
	http.HandleFunc("/", handlers.IndexHandler)

	http.ListenAndServe(":8090", nil)
}
