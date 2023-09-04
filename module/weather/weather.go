package weather

import (
	"encoding/json"
	"fmt"
	"github.com/vanisyd/tgbot/environment"
	"io"
	"log"
	"net/http"
)

func GetGeo(location string) Geolocation {
	url := buildUrl(ActionGeolocation{
		Query: location,
	})

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var locations []Geolocation
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}

	return locations[0]
}

func buildUrl(action Action) string {
	return fmt.Sprintf("%s%s&appid=%s", BaseUrl, action.GetQuery(), environment.Env.WeatherApiToken)
}
