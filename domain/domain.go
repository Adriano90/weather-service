package domain

type ForecastRepository interface {
	Forecast(latitude, longitude float64) (*Forecast, error)
}

type Forecast struct {
	City		city
	List	[]forecastData
}

type city struct {
	Id			int
	Name		string
	Coordinates	coordinates
	Country		string
}

type coordinates struct {
	Latitude	float64
	Longitude	float64
}

type forecastData struct {
	Dt			int
	Weather		[]weather
	Main		mainWeather
}

type weather struct {
	Id			int
	Main		string
	Description	string
	Icon		string
}

type mainWeather struct {
	Temp		float64
}