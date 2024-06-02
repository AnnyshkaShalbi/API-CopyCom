package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func accessDenied(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Access Denied")
	return
}

func api(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	usplit := strings.Split(r.URL.Path, "/api/")

	if len(usplit) == 1 {
		accessDenied(w, r)
		log.Println("тут 1")
		return
	} else {
		if _, ok := router[usplit[1]]; !ok {
			log.Println("тут 2")
			log.Println("Не найден event => ", usplit[1])
			accessDenied(w, r)
			return
		}
	}
	event := usplit[1]
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	router[event](w, r)
	return
}

func apiFile(w http.ResponseWriter, r *http.Request) {
	log.Println("apiFile")
	defer r.Body.Close()

	usplit := strings.Split(r.URL.Path, "/api/upload/")

	log.Printf("URL== %+v", usplit)
	event := usplit[1]
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	router[event](w, r)
	return
}
