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

type Image_params struct {
	Image string `json:"image"`
}

func Image(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params Image_params
	err := decoder.Decode(&params)

	log.Printf("file %+v", params.Image)

	if err != nil {
		log.Println("Error parse post News => ", err)
		fmt.Fprintf(w, "{}")
		return
	}

	mess := params.Image
	// mess += "👤 " + params.Name + "\n"
	// mess += "📞 " + params.Phone + "\n"
	// mess += "📬 " + params.Email + "\n"
	// mess += "📧 " + params.Comment

	fmt.Fprintf(w, "{}")

	chatID, _ := strconv.ParseInt(os.Getenv("TG_CHAT"), 10, 64)
	data := models.TgParams_sendMessage{
		Chat_id: chatID,
		Text:    mess,
	}

	out, _ := json.Marshal(data)
	telegram.SendImage(os.Getenv("TG_TOKEN"), out)

	return
}
