package domain

type ForecastRepository interface {
	Forecast(latitude, longitude float64) (*Forecast, error)
}

type Forecast struct {
	City city           `json:"city"`
	List []forecastData `json:"list"`
}

type city struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Coordinates coordinates `json:"coordinates"`
	Country     string      `json:"country"`
}

type coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type forecastData struct {
	Dt      int         `json:"dt"`
	Weather []weather   `json:"weather"`
	Main    mainWeather `json:"main"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type mainWeather struct {
	Temp float64 `json:"temp"`
}
