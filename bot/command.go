package bot

import (
	"fmt"
	"github.com/vanisyd/tgbot/database"
	"github.com/vanisyd/tgbot/module/weather"
	"time"
)

func Register([]string) string {
	database.AddUser(database.User{TgID: CurrentMSG.From.ID})
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
			locationName, ok := location.LocalNames["uk"]
			if ok {
				return fmt.Sprintf("У місті %s зараз %.1f°C", locationName, data.Current.Temp)
			} else {
				locationName, ok := location.LocalNames["en"]
				if ok {
					return fmt.Sprintf("У місті %s зараз %.1f°C", locationName, data.Current.Temp)
				} else {
					return weather.GenerateProverb()
				}
			}
		} else {
			return weather.GenerateProverb()
		}
	}

	return CurrentCMD.Signature()
}
