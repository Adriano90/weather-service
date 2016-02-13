package main

import (
	"fmt"
	"github.com/Adriano90/weather-service/interfaces"
	"github.com/Adriano90/weather-service/usecases"
	"github.com/drone/routes"
	"github.com/namsral/flag"
	"log"
	"net/http"
)

var (
	port             string
	openWeatherAppId string
)

func init() {
	flag.StringVar(&port, "port", "5005", "Server port number")
	flag.StringVar(&openWeatherAppId, "appid", "64b793d6792528f2b716206c1789ac82", "openweather.org app id")
}

func main() {
	flag.Parse()
	fmt.Println("Weather-service params, for more info use -help option:")
	fmt.Println("\t-Port: " + port)
	fmt.Println("\t-openweather app id: " + openWeatherAppId)

	fmt.Println("Bootstrapping system...")
	forecastInteractor := new(usecases.ForecastInteractor)
	forecastInteractor.ForecastRepository = interfaces.NewRestForecastRepo(&http.Client{}, openWeatherAppId)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.ForecastInteractor = forecastInteractor

	mux := routes.New()
	mux.Get("/forecast", webserviceHandler.GetForecast)
	http.Handle("/", mux)
	fmt.Println("Bootstrapping web service...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
