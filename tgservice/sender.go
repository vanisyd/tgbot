package tgservice

import (
	"fmt"
	"github.com/vanisyd/tgbot/config"
	"github.com/vanisyd/tgbot/environment"
)

func SetWebHookWithToken(token string) string {
	return BuildURLWithToken(config.RouteSetWebHook, token)
}

func SendMessage(token string) string {
	return BuildURLWithToken(config.RouteSendMsg, token)
}

func SetMenuButton(token string) string {
	return BuildURLWithToken(config.RouteSetMenuButton, token)
}

func BuildURL(action string) string {
	return fmt.Sprintf("%s%s/%s", config.BaseUrl, GetToken(), action)
}

func BuildURLWithToken(action string, token string) string {
	return fmt.Sprintf("%s%s/%s", config.BaseUrl, token, action)
}

func GetToken() string {
	return environment.Env.DevTgToken
}
