package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seu-usuario/weather-cep/internal/api/handlers"
)

func main() {
	http.HandleFunc("/weather/", handlers.HandleWeather)
	port := "8080"
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 