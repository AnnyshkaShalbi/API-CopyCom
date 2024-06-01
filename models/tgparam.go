package models

type TgParams_sendMessage struct {
	Chat_id int64  `json:"chat_id"`
	Text    string `json:"text"`
}
