package telegram

import (
	"copycoma-api/functions"
	"copycoma-api/models"
	"fmt"
	"log"
)

func GetUpdates(token string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", token)
	var answer = new(models.TgResp_Updates)
	err := functions.GetJson(url, answer)
	if err != nil {
		log.Println("Error GetUpdates", err)
	}
	log.Println("=2602e7=", answer.Result[0].My_chat_member.Chat.ID)
}

func SendMessage(token string, data []byte) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	var answer = new(interface{})
	if err := functions.PostJson(url, data, answer); err != nil {
		log.Println("Error SendMessage", err)
	}

	return
}
