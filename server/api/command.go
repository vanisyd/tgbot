package api

import (
	"fmt"
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/tgapi"
	"log"
	"os/exec"
)

func AttachBot(args []string) {
	token := args[0]
	hashID := args[1]

	payload := fmt.Sprintf("url=%s?hash_id=%s", HttpTunnel.URL(), hashID)
	cmd := exec.Command("curl", "-F", payload, tgapi.SetWebHookWithToken(token))

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	database.AddBot(database.Bot{
		Token:  token,
		HashID: hashID,
	})
	fmt.Printf("Bot attached: %s\n", hashID)
}
