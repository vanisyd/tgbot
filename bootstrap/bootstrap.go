package bootstrap

import (
	"context"
	"fmt"
	"github.com/vanisyd/tgbot/database"
	"github.com/vanisyd/tgbot/environment"
	"github.com/vanisyd/tgbot/tgapi"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"log"
	"os/exec"
)

func Init() ngrok.Tunnel {
	environment.Init()
	database.Init()
	return runServer(context.Background())
}

func runServer(ctx context.Context) ngrok.Tunnel {
	tunnel, err := ngrok.Listen(ctx, config.HTTPEndpoint(), ngrok.WithAuthtoken(environment.Env.NgrokAuthToken))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tunnel created:", tunnel.URL())
	cmd := exec.Command("curl", "-F", "url="+tunnel.URL(), tgapi.SetWebHook())

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started")
	return tunnel
}
