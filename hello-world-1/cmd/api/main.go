package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joaovitordesp/weather-cep/internal/api/handlers"
)

func main() {
	http.HandleFunc("/weather/", handlers.HandleWeather)
	fmt.Printf("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
