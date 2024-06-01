package events

import (
	"copycoma-api/models"
	"copycoma-api/telegram"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Message_params struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

func Message(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params Message_params
	err := decoder.Decode(&params)

	if err != nil {
		log.Println("Error parse post News => ", err)
		fmt.Fprintf(w, "{}")
		return
	}

	mess := "🎓 Новый заказ 🎓" + "\n\n"
	mess += "👤 " + params.Name + "\n"
	mess += "📞 " + params.Phone + "\n"
	mess += "📬 " + params.Email + "\n"
	mess += "📧 " + params.Comment

	fmt.Fprintf(w, "{}")

	chatID, _ := strconv.ParseInt(os.Getenv("TG_CHAT"), 10, 64)
	data := models.TgParams_sendMessage{
		Chat_id: chatID,
		Text:    mess,
	}

	out, _ := json.Marshal(data)
	telegram.SendMessage(os.Getenv("TG_TOKEN"), out)

	return
}
