package bootstrap

import (
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/environment"
	"github.com/vanisyd/tgbot/server"
	"github.com/vanisyd/tgbot/server/api"
)

func Init() {
	environment.Init()
	database.Init(environment.Env.DBUri, environment.Env.DBName)
	api.Init()
	server.Init()
}
