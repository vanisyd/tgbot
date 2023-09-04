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
		location := weather.GetGeo(params[0])
		fmt.Println(location)
	}
	return "Погода"
}
