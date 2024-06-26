package bot

import (
	"fmt"
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/module/weather"
	"time"
)

func Register([]string) (string, interface{}) {
	database.AddUser(database.User{TgID: CurrentMSG.From.ID})
	return fmt.Sprintf("Welcome, %s!", CurrentMSG.From.Username), nil
}

func Now([]string) (string, interface{}) {
	return fmt.Sprintf("Дата: %s", time.Now().String()), nil
}

func Weather(params []string) (string, interface{}) {
	if len(params) > 0 {
		locations := weather.GetGeo(params[0])
		if len(locations) > 0 {
			user := GetCurrentDBUser()
			database.AddAction(database.Action{
				UserId: user.ID,
				Data:   locations,
			})

			location := locations[0]
			data := weather.GetWeather(location)
			locationName, ok := location.LocalNames["uk"]
			if ok {
				return fmt.Sprintf("У місті %s зараз %.1f°C", locationName, data.Current.Temp), nil
			} else {
				locationName, ok := location.LocalNames["en"]
				if ok {
					return fmt.Sprintf("У місті %s зараз %.1f°C", locationName, data.Current.Temp), nil
				} else {
					return weather.GenerateProverb(), nil
				}
			}
		} else {
			return weather.GenerateProverb(), nil
		}
	}

	return CurrentCMD.Signature(), nil
}
