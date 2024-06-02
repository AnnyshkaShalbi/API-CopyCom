package server

import (
	"copycoma-api/events"
	"log"
	"net/http"
)

var router = map[string]func(w http.ResponseWriter, r *http.Request){
	"Message":    events.Message,
	"FileUpload": events.FileUpload,
}

func Start(host string) {
	http.HandleFunc("/", api)
	http.HandleFunc("/api/upload/", apiFile)
	log.Fatal("HTTP server error: ", http.ListenAndServe(host, nil))
}
