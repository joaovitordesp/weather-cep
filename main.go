package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type Temperature struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/weather/", handleWeather)
	port := "8080"
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	cep := r.URL.Path[len("/weather/"):]
	
	if !isValidCEP(cep) {
		respondWithError(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	city, err := getLocationByCEP(cep)
	if err != nil {
		if err.Error() == "can not find zipcode" {
			respondWithError(w, http.StatusNotFound, "can not find zipcode")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Erro interno do servidor")
		return
	}

	tempC, err := getTemperature(city)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro ao obter temperatura")
		return
	}

	temps := convertTemperatures(tempC)
	respondWithJSON(w, http.StatusOK, temps)
}

func isValidCEP(cep string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Message: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
} 