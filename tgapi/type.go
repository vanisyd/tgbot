package tgapi

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
