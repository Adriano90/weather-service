package interfaces

import (
	"github.com/Adriano90/weather-service/domain"
	"github.com/drone/routes"
	"log"
	"net/http"
	"strconv"
)

type ForecastInteractor interface {
	FindByCoordinates(latitude, longitude float64) (*domain.Forecast, error)
}

type WebserviceHandler struct {
	ForecastInteractor ForecastInteractor
}

func (handler WebserviceHandler) GetForecast(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	latitude, err := strconv.ParseFloat(params.Get("latitude"), 64)

	if err != nil {
		log.Printf("Error parsing latitude: %s", err.Error())
		http.Error(w, "Wrong latitude value", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(params.Get("longitude"), 64)

	if err != nil {
		log.Printf("Error parsing longitude: %s", err.Error())
		http.Error(w, "Wrong longitude value", http.StatusBadRequest)
		return
	}

	forecast, err := handler.ForecastInteractor.FindByCoordinates(latitude, longitude)

	if err != nil {
		log.Printf("Error retrieving forecast: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	routes.ServeFormatted(w, r, forecast)
}
