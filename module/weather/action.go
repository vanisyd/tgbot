package weather

import "fmt"

// ActionGeolocation Locations API
type ActionGeolocation struct {
	Query string
	Limit int
}

func (act ActionGeolocation) GetQuery() string {
	limit := act.Limit
	if limit == 0 {
		limit = 20
	}

	return fmt.Sprintf("%s?q=%s&limit=%d", GeoUrl, act.Query, limit)
}

// ActionWeather Weather API
type ActionWeather struct {
	Lat float32
	Lon float32
}

func (act ActionWeather) GetQuery() string {
	return fmt.Sprintf("%s?lat=%f&lon=%f&units=metric", WeatherUrl, act.Lat, act.Lon)
}
