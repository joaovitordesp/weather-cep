package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaovitordesp/weather-cep/internal/api/handlers"
)

func TestHandleWeather(t *testing.T) {
	tests := []struct {
		name           string
		cep            string
		expectedStatus int
	}{
		{"CEP Inválido", "123", 422},
		{"CEP Não Encontrado", "99999999", 404},
		{"CEP Válido", "01001000", 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/weather/"+tt.cep, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.HandleWeather)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
} 