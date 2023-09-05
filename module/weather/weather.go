package weather

import (
	"encoding/json"
	"fmt"
	"github.com/vanisyd/tgbot/environment"
	"io"
	"log"
	"net/http"
)

func GetGeo(location string) (locations []Geolocation) {
	fetchData(ActionGeolocation{
		Query: location,
	}, &locations)

	return
}

func GetWeather(location Geolocation) (weather Weather) {
	fetchData(ActionWeather{
		Lon: location.Lon,
		Lat: location.Lat,
	}, &weather)

	return
}

func buildUrl(action Action) string {
	return fmt.Sprintf("%s%s&appid=%s", BaseUrl, action.GetQuery(), environment.Env.WeatherApiToken)
}

func fetchData(action Action, dest any) {
	url := buildUrl(action)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, dest)
	if err != nil {
		log.Fatal(err)
	}
}
