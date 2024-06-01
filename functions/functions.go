package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Println("=f97461=", err)
		return err
	}
	if r.StatusCode != 200 {
		log.Println("=15e3ab= err status ", r.StatusCode)
		return errors.New("Status not 200")
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func PostJson(url string, data []byte, target interface{}) error {
	r, err := myClient.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("=f97461=", err)
		return err
	}
	if r.StatusCode != 200 {
		log.Println("=15e3ab= err status ", r.StatusCode)
		var answ = new(interface{})
		err := json.NewDecoder(r.Body).Decode(answ)
		dat, _ := json.Marshal(answ)
		log.Println("=328bf3=", err, string(dat))
		return errors.New("Status not 200")
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
