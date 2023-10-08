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

func Handler(res http.ResponseWriter, req *http.Request) {
	body := ResponseBody{}
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		fmt.Println("[JSON]", err)
		return
	}

	bot.CurrentMSG = body.Message

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

func sendMsg(chatID int64, message string, markup interface{}) { // TODO: refactor methods
	body := RequestBody{
		ChatID: chatID,
		Text:   message,
	}

	if markup != nil {
		body.ReplyMarkup = ReplyMarkupData{
			InlineKeyboard: markup,
		}
	}

	reqBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Send message")
	res, err := http.Post(tgapi.SendMessage(), "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		var errMsg interface{}
		if err = json.NewDecoder(res.Body).Decode(&errMsg); err != nil {
			log.Fatal(err)
		}
		log.Fatal(errMsg)
		//log.Fatal("Unexpected status", res.Status)
	}
}

func Welcome(chatID int64, user tgapi.User) {
	sendMsg(chatID, fmt.Sprintf("Hi, %s! Welcome to my bot", user.Username), nil)
}
