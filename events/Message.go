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
	Name              string `json:"name"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Comment           string `json:"comment"`
	Color             bool   `json:"color"`
	TitleCover        string `json:"titleCover"`
	TitleLogo         string `json:"titleLogo"`
	TotalPages        int    `json:"totalPages"`
	CountBlackPages   int    `json:"countBlackPages"`
	CountColorPages   int    `json:"countColorPages"`
	ColoredPages      []int  `json:"coloredPages"`
	Price             int    `json:"price"`
	PocketForReview   bool   `json:"pocketForReview"`
	PocketDiskCD      bool   `json:"pocketDiskCD"`
	PlasticFileBefore struct {
		Active   bool `json:"active"`
		Quantity int  `json:"quantity"`
	} `json:"plasticFileBefore"`
	PlasticFileAfter struct {
		Active   bool `json:"active"`
		Quantity int  `json:"quantity"`
	} `json:"plasticFileAfter"`
	PlastikFileInTheEnd struct {
		Active   bool `json:"active"`
		Quantity int  `json:"quantity"`
	} `json:"plastikFileInTheEnd"`
}

func Message(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params Message_params
	err := decoder.Decode(&params)

	if err != nil {
		log.Println("Error parse request copycom => ", err)
		fmt.Fprintf(w, "{}")
		return
	}

	mess := "🎓 Новый заказ 🎓" + "\n\n"
	mess += "👤Имя: " + params.Name + "\n"
	mess += "📞Телефон: " + params.Phone + "\n"
	mess += "📬Email: " + params.Email + "\n"
	mess += "📧Комментарий: " + params.Comment + "\n"
	mess += "📑Всего страниц: " + fmt.Sprint(params.TotalPages) + "\n"
	mess += "📑Количество страниц ч/б: " + fmt.Sprint(params.CountBlackPages) + "\n"
	mess += "📑Количество цветных страниц: " + fmt.Sprint(params.CountColorPages) + "\n"

	if len(params.ColoredPages) == 0 {
		mess += "📑Цветные страницы не указаны!" + "\n"
	} else {
		mess += "\n" + "Цветные страницы: "

		for _, number := range params.ColoredPages {
			mess += fmt.Sprint(number) + ","
		}

		mess += "\n"
	}

	if params.Color {
		mess += "\n" + "🟥 Красная обложка 🟥" + "\n"
	} else {
		mess += "\n" + "🟦 Синяя обложка 🟦" + "\n"
	}

	mess += "Заголовок обложки: " + params.TitleCover + "\n"
	mess += "Заголовок логотипа: " + params.TitleLogo + "\n"

	if params.PocketForReview {
		mess += "✅ Вклеить карман для рецензии ✅" + "\n"
	} else {
		mess += "❌ Не вклеивать карман для рецензии ❌" + "\n"
	}

	if params.PocketDiskCD {
		mess += "\n" + "✅ Вклеить карман для CD диска ✅" + "\n"
	} else {
		mess += "\n" + "❌ Не вклеивать карман для CD диска ❌" + "\n"
	}

	if params.PlasticFileBefore.Active {
		mess += "\n" + "💿 Добавить пластиковый файл перед титулом" + "\n"
		mess += "Количество: " + fmt.Sprint(params.PlasticFileBefore.Quantity) + "\n"
	}

	if params.PlasticFileAfter.Active {
		mess += "\n" + "💿 Добавить пластиковый файл после титулом" + "\n"
		mess += "Количество: " + fmt.Sprint(params.PlasticFileAfter.Quantity) + "\n"
	}

	if params.PlastikFileInTheEnd.Active {
		mess += "\n" + "💿 Добавить пластиковый файл в конце работы" + "\n"
		mess += "Количество: " + fmt.Sprint(params.PlastikFileInTheEnd.Quantity) + "\n"
	}

	mess += "💰ЦЕНА💰: " + fmt.Sprint(params.Price) + " ₽" + "\n"

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
