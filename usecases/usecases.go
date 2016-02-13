package usecases

import (
	"github.com/Adriano90/weather-service/domain"
	"log"
)

type ForecastInteractor struct {
	ForecastRepository domain.ForecastRepository
}

func (interactor *ForecastInteractor) FindByCoordinates(latitude, longitude float64) (*domain.Forecast, error) {

	log.Printf("usecases FindByCoordinates %f %f", latitude, longitude)
	return interactor.ForecastRepository.Forecast(latitude, longitude)
}
