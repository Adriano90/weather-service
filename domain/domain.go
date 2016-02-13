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
	Country     string      `json:"country"`
	Coordinates coordinates `json:"coord"`
}

type coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type forecastData struct {
	Dt       int         `json:"dt"`
	Weather  []weather   `json:"weather"`
	Temp     temperature `json:"temp"`
	Pressure float64     `json:"pressure"`
	Humidity float64     `json:"humidity"`
	Speed    float64     `json:"speed"`
	Deg      float64     `json:"deg"`
	Clouds   float64     `json:"clouds"`
	Rain     float64     `json:"rain,omitempty"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type temperature struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}
