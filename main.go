package main

import (
	"github.com/drone/routes"
	"net/http"
	"github.com/Adriano90/weather-service/interfaces"
	"github.com/Adriano90/weather-service/usecases"
	"log"
)

func main() {

	restHandler := &http.Client{}
	handlers := make(map[string] interfaces.RestHandler)
	handlers["RestForecastRepo"] = restHandler

	forecastInteractor := new(usecases.ForecastInteractor)
	forecastInteractor.ForecastRepository = interfaces.NewRestForecastRepo(handlers)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.ForecastInteractor = forecastInteractor

	mux := routes.New()
	mux.Get("/forecast/:latitude/:longitude", webserviceHandler.GetForecast)
	http.Handle("/", mux)
	log.Println("Bootstrapping weather-service...")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
