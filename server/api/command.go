package api

import (
	"fmt"
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/server"
)

func AttachBot(args []string) {
	token := args[0]
	hashID := args[1]

	server.InitBot(token, hashID)
	database.AddBot(database.Bot{
		Token:  token,
		HashID: hashID,
	})
	fmt.Printf("Bot attached: %s\n", hashID)
}
