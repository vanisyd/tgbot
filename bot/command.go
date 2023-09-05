package bot

import (
	"fmt"
	"github.com/vanisyd/tgbot/module/weather"
	"time"
)

func Register([]string) string {
	return fmt.Sprintf("Welcome, %s!", CurrentMSG.From.Username)
}

func Now([]string) string {
	return fmt.Sprintf("Дата: %s", time.Now().String())
}

func Weather(params []string) string {
	if len(params) > 0 {
		locations := weather.GetGeo(params[0])
		if len(locations) > 0 {
			location := locations[0]
			data := weather.GetWeather(location)
			return fmt.Sprintf("У місті %s зараз %.1f°C", location.LocalNames["uk"], data.Current.Temp)
		} else {
			return "Погоду не знайдено"
		}
	}

	return CurrentCMD.Signature()
}
