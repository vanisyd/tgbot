package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/vanisyd/tgapi/kind"
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/bot"
	"github.com/vanisyd/tgbot/environment"
	"github.com/vanisyd/tgbot/tgservice"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"log"
	"net/http"
	"os/exec"
)

var Tunnel ngrok.Tunnel
var CurrentBot database.Bot

func Init() {
	RunServer(context.Background())
	initBots()
	err := http.Serve(Tunnel, http.HandlerFunc(HttpHandler))
	if err != nil {
		log.Fatal(err)
	}
}

func RunServer(ctx context.Context) ngrok.Tunnel {
	tunnel, err := ngrok.Listen(ctx, config.HTTPEndpoint(), ngrok.WithAuthtoken(environment.Env.NgrokAuthToken))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tunnel created:", tunnel.URL())
	Tunnel = tunnel

	return tunnel
}

func HttpHandler(_ http.ResponseWriter, req *http.Request) {
	hashID := req.URL.Query().Get("hash_id")
	fmt.Printf("Request from %s\n", hashID)
	dbBot := database.FindBot(hashID)
	if botObj, ok := dbBot.(database.Bot); ok {
		CurrentBot = botObj

		body := kind.ResponseBody{}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			fmt.Println("[JSON]", err)
			return
		}

		if len(body.Message.Text) > 0 && body.Message.Text[0] == '/' {
			command, err := bot.GetCMD(body.Message.Text)
			if err != nil {
				sendMsg(body.Message.Chat.ID, "Unknown command", nil)
				log.Printf("[ERROR] Code: 400 Bot hash: %s. Unknown command %s\n", hashID, body.Message.Text)
			} else {
				response, markup := command.Handler(bot.GetParams(body.Message.Text))
				sendMsg(body.Message.Chat.ID, response, markup)
			}
		} else {
			Welcome(body.Message.Chat.ID, body.Message.From)
		}
	} else {
		fmt.Println("Unknown request")
	}
}

func sendMsg(chatID int64, message string, markup interface{}) {
	body := kind.SendMessageBody{
		ChatID: chatID,
		Text:   message,
	}

	if markup != nil {
		body.ReplyMarkup = kind.ReplyMarkupData{
			InlineKeyboard: markup,
		}
	}

	sendRequest(body, tgservice.SendMessage(CurrentBot.Token))
}

func setMenuBtn(chatID int64) {
	body := kind.SetChatMenuButtonBody{
		ChatID: chatID,
		MenuButton: kind.MenuButtonWebApp{
			Type: kind.MenuButtonTypeWebApp,
			Text: "Shop",
			WebApp: kind.WebAppInfo{
				URL: bot.BuildWebAppURL(bot.MenuButtonRoute),
			},
		},
	}

	sendRequest(body, tgservice.SetMenuButton(CurrentBot.Token))
}

func sendRequest(reqBody any, action string) {
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(action, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		var errMsg interface{}
		if err = json.NewDecoder(res.Body).Decode(&errMsg); err != nil {
			log.Fatal(err)
		}
		log.Fatal(errMsg)
	}
}

func Welcome(chatID int64, user kind.User) {
	sendMsg(chatID, fmt.Sprintf("Hi, %s! Welcome to my bot", user.Username), nil)
}

func InitBot(token string, hashID string) {
	payload := fmt.Sprintf("url=%s?hash_id=%s", Tunnel.URL(), hashID)
	cmd := exec.Command("curl", "-F", payload, tgservice.SetWebHookWithToken(token))

	if err := cmd.Run(); err != nil {
		log.Println("[TCP.Error] " + err.Error())
		return
	}

	log.Printf("Bot %s loaded", hashID)
}

func initBots() {
	bots := database.GetBots()
	if botsData, ok := bots.([]database.Bot); ok {
		for _, botData := range botsData {
			InitBot(botData.Token, botData.HashID)
		}
	}
}
