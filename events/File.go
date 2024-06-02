package events

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// type File interface {
// 	io.Reader
// 	io.ReaderAt
// 	io.Seeker
// 	io.Closer
// }

// type FileHeader struct {
// 	Filename string
// 	Header   textproto.MIMEHeader
// 	Size     int64
// }

// type Form struct {
// 	Value map[string][]string
// 	File  map[string][]*FileHeader
// }

// Part представляет собой одну деталь в составном теле
// type Part struct {
// 	Header textproto.MIMEHeader
// }

// Reader — это итератор по частям в составном теле MIME. Базовый синтаксический анализатор Reader потребляет входные данные по мере необходимости. Ищущий не поддерживается.
// type Reader struct {
// }

// Open открывает и возвращает File, связанный с FileHeader.
// func (fh *FileHeader) Open() (File, error)

// RemoveAll удаляет все временные файлы, связанные с формой.
// func (f *Form) RemoveAll() error

// func (p *Part) Close() error

/*
FileName возвращает параметр filename объекта Content-Disposition компонента заголовок.
Если имя файла не пустое, то оно передается через путь к файлу.
Base (который является зависит от платформы) перед возвратом.
*/
// func (p *Part) FileName()

/*
FormName возвращает параметр name, если p имеет Content-Disposition типа "form-data".
В противном случае возвращается пустая строка.
*/
// func (p *Part) FormName() string

// Read читает тело части, после ее заголовков и перед Начинается следующая часть (если таковая имеется).
// func (p *Part) Read(d []byte) (n int, err error)

// NewReader создает новый составной считыватель Reader из r с помощью метода заданной границей MIME.
// func NewReader(r io.Reader, boundary string) *Reader

func FileUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("FileUpload")

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
	log.Println("sendFileToTelegram")
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
