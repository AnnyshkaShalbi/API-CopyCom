package events

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 10 << 20 соответствует 10MB, максимальный размер файла
	r.ParseMultipartForm(10 << 20)

	// Получаем файл из запроса
	file, _, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем файл
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendFileToTelegram(fileBytes, "PDF File", os.Getenv("TG_CHAT"), os.Getenv("TG_TOKEN"))
	// Здесь можно сохранить файл на диск или выполнить другие действия

	// Отправляем ответ клиенту
	w.Write([]byte("Файл успешно получен"))
}

func sendFileToTelegram(fileBytes []byte, fileName string, chatID string, botToken string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", botToken)

	// Создаем буфер для записи многочастного сообщения
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Добавляем файл
	part, err := writer.CreateFormFile("document", filepath.Base(fileName))
	if err != nil {
		return err
	}
	part.Write(fileBytes)

	// Добавляем дополнительные поля, если необходимо
	_ = writer.WriteField("chat_id", chatID)

	// Завершаем запись тела запроса
	err = writer.Close()
	if err != nil {
		return err
	}

	// Создаем запрос
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return err
	}

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем ответ от сервера
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send document: %s", resp.Status)
	}

	return nil
}
