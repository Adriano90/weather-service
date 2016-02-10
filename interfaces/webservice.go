package interfaces

import (
	"encoding/json"
	"github.com/Adriano90/weather-service/domain"
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
	latitude, _ := strconv.ParseFloat(params.Get(":latitude"), 64)
	longitude, _ := strconv.ParseFloat(params.Get(":longitude"), 64)
	forecast, _ := handler.ForecastInteractor.FindByCoordinates(latitude, longitude)
	js, err := json.Marshal(forecast)

	if err != nil {
		log.Printf("Error retrieving forecast: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
