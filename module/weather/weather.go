package weather

import (
	"encoding/json"
	"fmt"
	"github.com/vanisyd/tgbot/environment"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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

func GenerateProverb() string {
	file, err := os.ReadFile("sentences.source")
	if err != nil {
		log.Fatal(err)
	}
	proverbs := strings.Split(string(file), "\n")
	line := rand.Intn(len(proverbs))

	return proverbs[line]
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
