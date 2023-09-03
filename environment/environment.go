package environment

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Env EnvVars

func Init() {
	fmt.Println("Loading environment")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env file")
	}

	Env.NgrokAuthToken = os.Getenv("NGROK_AUTH_TOKEN")
	Env.DevTgToken = os.Getenv("DEV_TG_TOKEN")
}
