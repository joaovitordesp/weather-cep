package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/seu-usuario/weather-cep/internal/models"
	"github.com/seu-usuario/weather-cep/internal/services/viacep"
	"github.com/seu-usuario/weather-cep/internal/services/weather"
)

func HandleWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	cep := r.URL.Path[len("/weather/"):]
	
	if !isValidCEP(cep) {
		respondWithError(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	city, err := viacep.GetLocationByCEP(cep)
	if err != nil {
		if err.Error() == "can not find zipcode" {
			respondWithError(w, http.StatusNotFound, "can not find zipcode")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Erro interno do servidor")
		return
	}

	tempC, err := weather.GetTemperature(city)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro ao obter temperatura")
		return
	}

	temps := weather.ConvertTemperatures(tempC)
	respondWithJSON(w, http.StatusOK, temps)
}

func isValidCEP(cep string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.ErrorResponse{Message: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
} 