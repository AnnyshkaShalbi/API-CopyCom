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

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func api(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setupCORS(&w)
	if r.Method == http.MethodOptions {
		w.Write([]byte("OK"))
		return
	}

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

	setupCORS(&w)
	if r.Method == http.MethodOptions {
		w.Write([]byte("OK"))
		return
	}

	usplit := strings.Split(r.URL.Path, "/api/upload/")

	log.Printf("URL== %+v", usplit)
	event := usplit[1]
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	router[event](w, r)
	return
}
