package main

import (
	"copycoma-api/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	host := "127.0.0.1:80"
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Файл ENV не найден")
		host = ":80"
	}

	server.Start(host)
}
