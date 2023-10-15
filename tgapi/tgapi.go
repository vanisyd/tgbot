package tgapi

import (
	"fmt"
	"github.com/vanisyd/tgbot/config"
	"github.com/vanisyd/tgbot/environment"
)

// TODO: refactor duplicating code

func SetWebHookWithToken(token string) string {
	return BuildURLWithToken(RouteSetWebHook, token)
}

func SendMessage(token string) string {
	return BuildURLWithToken(RouteSendMsg, token)
}

func SetMenuButton(token string) string {
	return BuildURLWithToken(RouteSetMenuButton, token)
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
