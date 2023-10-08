package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vanisyd/tgbot/bot"
	"github.com/vanisyd/tgbot/tgapi"
	"log"
	"net/http"
)

func Handler(_ http.ResponseWriter, req *http.Request) {
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
}

func sendMsg(chatID int64, message string, markup interface{}) { // TODO: refactor method
	body := tgapi.SendMessageBody{
		ChatID: chatID,
		Text:   message,
	}

	if markup != nil {
		body.ReplyMarkup = tgapi.ReplyMarkupData{
			InlineKeyboard: markup,
		}
	}

	sendRequest(body, tgapi.SendMessage())
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

	sendRequest(body, tgapi.SetMenuButton())
}

func sendRequest(reqBody any, action string) {
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Request sent: %s", action)
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
