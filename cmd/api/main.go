package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joaovitordesp/weather-cep/internal/api/handlers"
)

func main() {
	http.HandleFunc("/weather/", handlers.HandleWeather)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
