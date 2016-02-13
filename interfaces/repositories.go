package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/Adriano90/weather-service/domain"
	"io/ioutil"
	"log"
	"net/http"
)

type RestHandler interface {
	Do(request *http.Request) (*http.Response, error)
}

type RestForecastRepo struct {
	appId       string
	restHandler RestHandler
}

func NewRestForecastRepo(restHandler RestHandler, appId string) *RestForecastRepo {
	repo := &RestForecastRepo{
		appId:       appId,
		restHandler: restHandler,
	}
	return repo
}

func (restForecastRepo *RestForecastRepo) Forecast(latitude, longitude float64) (*domain.Forecast, error) {

	const (
		maxForecast    = 16
		openWeatherUrl = "http://api.openweathermap.org/data/2.5/forecast/daily?lat=%f&lon=%f&appid=%s&cnt=%d"
	)

	url := fmt.Sprintf(
		openWeatherUrl,
		latitude,
		longitude,
		restForecastRepo.appId,
		maxForecast,
	)

	log.Printf("Request URL: %s", url)

	req, requestError := http.NewRequest("GET", url, nil)

	if requestError != nil {
		log.Fatal("Error creating request: %s", requestError.Error())
		return nil, requestError
	}

	resp, error := restForecastRepo.restHandler.Do(req)

	if error != nil {
		log.Fatal("Error invoking openweather api: %s", error.Error())
		return nil, error
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	forecast := &domain.Forecast{}
	errorParsing := json.Unmarshal(body, forecast)

	if errorParsing != nil {
		log.Fatal("Error parsing openweather api response: %s", errorParsing.Error())
	}

	return forecast, errorParsing
}
