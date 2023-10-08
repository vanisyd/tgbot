package bot

import (
	"fmt"
	"github.com/vanisyd/tgbot/environment"
	"strings"
)

const OptionsRoute string = "options"
const MenuButtonRoute string = "shop"

func BuildWebAppURL(route string) string {
	webAppUrl := environment.Env.WebAppURL
	if webAppUrl[len(webAppUrl)-1] != '/' {
		webAppUrl = webAppUrl + "/"
	}

	user := GetCurrentDBUser()
	routeUrl := route
	if strings.Contains(routeUrl, "?") {
		routeUrl = routeUrl + "&"
	} else {
		if routeUrl[len(routeUrl)-1] == '/' {
			routeUrl = routeUrl[:len(routeUrl)-2]
		}
		routeUrl = routeUrl + "?"
	}
	routeUrl = fmt.Sprintf("%suser_id=%s", routeUrl, user.ID.String())

	return fmt.Sprintf("%s%s", environment.Env.WebAppURL, routeUrl)
}
