package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	database "github.com/vanisyd/tgbot-db"
	"github.com/vanisyd/tgbot/bot"
	"github.com/vanisyd/tgbot/environment"
	"github.com/vanisyd/tgbot/server/api"
	"github.com/vanisyd/tgbot/tgapi"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"log"
	"net/http"
)

var Tunnel ngrok.Tunnel
var CurrentBot database.Bot

func Init() {
	RunServer(context.Background())
	api.Init(Tunnel)
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

		body := tgapi.ResponseBody{}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			fmt.Println("[JSON]", err)
			return
		}

		bot.CurrentMSG = body.Message

		setMenuBtn(body.Message.Chat.ID)

		if len(body.Message.Text) > 0 && body.Message.Text[0] == '/' {
			command, err := bot.GetCMD(body.Message.Text)
			if err != nil {
				sendMsg(body.Message.Chat.ID, "Неправильна команда", nil)
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
	body := tgapi.SendMessageBody{
		ChatID: chatID,
		Text:   message,
	}

	if markup != nil {
		body.ReplyMarkup = tgapi.ReplyMarkupData{
			InlineKeyboard: markup,
		}
	}

	sendRequest(body, tgapi.SendMessage(CurrentBot.Token))
}

func setMenuBtn(chatID int64) {
	body := tgapi.SetChatMenuButtonBody{
		ChatID: chatID,
		MenuButton: tgapi.MenuButtonWebApp{
			Type: tgapi.MenuButtonTypeWebApp,
			Text: "Shop",
			WebApp: tgapi.WebAppInfo{
				URL: bot.BuildWebAppURL(bot.MenuButtonRoute),
			},
		},
	}

	sendRequest(body, tgapi.SetMenuButton(CurrentBot.Token))
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

func Welcome(chatID int64, user tgapi.User) {
	sendMsg(chatID, fmt.Sprintf("Hi, %s! Welcome to my bot", user.Username), nil)
}
