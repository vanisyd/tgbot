package server

import "github.com/vanisyd/tgbot/tgapi"

type ResponseBody struct {
	Message tgapi.Message `json:"message"`
}

type RequestBody struct { // TODO: change to 'sendMessage interface'
	ChatID      int64           `json:"chat_id"`
	Text        string          `json:"text"`
	ReplyMarkup ReplyMarkupData `json:"reply_markup,omitempty"`
}

type ReplyMarkupData struct {
	InlineKeyboard interface{} `json:"inline_keyboard,omitempty"`
}
