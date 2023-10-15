package bootstrap

import (
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/environment"
	"github.com/vanisyd/tgbot/server"
)

func Init() {
	environment.Init()
	database.Init(environment.Env.DBUri, environment.Env.DBName)
	server.Init()
}
