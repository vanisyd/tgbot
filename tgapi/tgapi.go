package tgapi

import (
	"fmt"
	"github.com/vanisyd/tgbot/config"
	"github.com/vanisyd/tgbot/environment"
)

func SetWebHook() string {
	return BuildURL(RouteSetWebHook)
}

func SendMessage() string {
	return BuildURL(RouteSendMsg)
}

func BuildURL(action string) string {
	return fmt.Sprintf("%s%s/%s", config.BaseUrl, GetToken(), action)
}

func GetToken() string {
	return environment.Env.DevTgToken
}
