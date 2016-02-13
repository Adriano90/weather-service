package main

import (
	"github.com/Adriano90/weather-service/Godeps/_workspace/src/github.com/drone/routes"
	"github.com/Adriano90/weather-service/cmd/weather-service/interfaces"
	"github.com/Adriano90/weather-service/cmd/weather-service/usecases"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	restHandler := &http.Client{}
	handlers := make(map[string]interfaces.RestHandler)
	handlers["RestForecastRepo"] = restHandler

	forecastInteractor := new(usecases.ForecastInteractor)
	forecastInteractor.ForecastRepository = interfaces.NewRestForecastRepo(handlers)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.ForecastInteractor = forecastInteractor

	mux := routes.New()
	mux.Get("/forecast", webserviceHandler.GetForecast)
	http.Handle("/", mux)
	log.Println("Bootstrapping weather-service...")
	log.Fatal(http.ListenAndServe(port, nil))
}
