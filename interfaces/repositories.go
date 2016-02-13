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

type RestRepo struct {
	restHandler RestHandler
}

type RestForecastRepo struct {
	RestHandlers map[string]RestHandler
	RestHandler  RestHandler
}

func NewRestForecastRepo(restHandlers map[string]RestHandler) *RestForecastRepo {
	repo := new(RestForecastRepo)
	repo.RestHandlers = restHandlers
	repo.RestHandler = restHandlers["RestForecastRepo"]
	return repo
}

func (restForecastRepo *RestForecastRepo) Forecast(latitude, longitude float64) (*domain.Forecast, error) {

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&cnt=%d", latitude, longitude, "64b793d6792528f2b716206c1789ac82", 2)
	log.Printf("Request URL: %s", url)
	forecast := new(domain.Forecast)

	req, requestError := http.NewRequest("GET", url, nil)

	if requestError != nil {
		log.Fatal("Error creating request: %s", requestError.Error())
		return nil, requestError
	}

	resp, error := restForecastRepo.RestHandler.Do(req)

	if error != nil {
		log.Fatal("Error invoking openweather api: %s", error.Error())
		return nil, error
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	errorParsing := json.Unmarshal(body, forecast)

	if errorParsing != nil {
		log.Fatal("Error parsing openweather api response: %s", errorParsing.Error())
	}

	return forecast, errorParsing
}
