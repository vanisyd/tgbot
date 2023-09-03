package server

import "github.com/vanisyd/tgbot/tgapi"

type ResponseBody struct {
	Message tgapi.Message `json:"message"`
}

type RequestBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
