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

type Message_params_service struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	TotalPages int    `json:"totalPages"`
	Action     string `json:"action"`
}

func UploadService(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params Message_params_service
	err := decoder.Decode(&params)

	if err != nil {
		log.Println("Error parse request copycom => ", err)
		fmt.Fprintf(w, "{}")
		return
	}
	fmt.Println(params.Name)

	mess := "🎓 Новый заказ 🎓" + "\n"

	if params.Action == "projectdoc" {
		mess = "🗃 ПРОЕКТНАЯ ДОКУМЕНТАЦИЯ 🗃" + "\n\n"
	}

	if params.Action == "copydoc" {
		mess = "🖨 ПЕЧАТЬ ДОКУМЕНТОВ 🖨" + "\n\n"
	}

	if params.Action == "drawings" {
		mess = "🖨 ПЕЧАТЬ ЧЕРТЕЖЕЙ 🖨" + "\n\n"
	}

	if params.Action == "presentations" {
		mess = "🖨 ПЕЧАТЬ ПРЕЗЕНТАЦИЙ 🖨" + "\n\n"
	}

	if params.Action == "patterns" {
		mess = "🖨 ПЕЧАТЬ ЛЕКАЛ И ВЫКРОЕК 🖨" + "\n\n"
	}

	if params.Action == "scanning" {
		mess = "🖨 СКАНИРОВАНИЕ ДОКУМЕНТОВ 🖨" + "\n\n"
	}

	if params.Action == "hardcover" {
		mess = "🖨 ТВЁРДЫЙ ПЕРЕПЛЁТ ДИПЛОМОВ 🖨" + "\n\n"
	}

	if params.Action == "brochure" {
		mess = "🖨 БРОШЮРОВКА НА ПЛАСТИКОВУЮ ПРУЖИНУ 🖨" + "\n\n"
	}

	mess += "📂 Имя файла: " + params.Name + "\n"
	mess += "📞 Телефон клиента: " + params.Phone + "\n"

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
