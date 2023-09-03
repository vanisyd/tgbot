package main

import (
	"fmt"
	"github.com/vanisyd/tgbot/bootstrap"
	"github.com/vanisyd/tgbot/server"
	"log"
	"net/http"
)

func main() {
	tunnel := bootstrap.Init()

	err := http.Serve(tunnel, http.HandlerFunc(server.Handler))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is ready")
}
