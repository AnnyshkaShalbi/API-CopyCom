package server

import (
	"copycoma-api/events"
	"log"
	"net/http"
)

var router = map[string]func(w http.ResponseWriter, r *http.Request){
	"Message": events.Message,
	"Image":   events.Image,
}

func Start(host string) {
	http.HandleFunc("/", api)
	log.Fatal("HTTP server error: ", http.ListenAndServe(host, nil))
}
