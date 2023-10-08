package tgapi

type ResponseBody struct {
	Message Message `json:"message"`
}

type SendMessageBody struct {
	ChatID      int64       `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SetChatMenuButtonBody struct {
	ChatID     int64       `json:"chat_id"`
	MenuButton interface{} `json:"menu_button"`
}

type ReplyMarkupData struct {
	InlineKeyboard interface{} `json:"inline_keyboard,omitempty"`
}

type Message struct {
	Text string `json:"text"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firs_name"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type InlineKeyboardButton struct {
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}

type MenuButtonWebApp struct {
	Type   string     `json:"type"`
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}
