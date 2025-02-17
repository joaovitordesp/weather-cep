package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	
	"github.com/seu-usuario/weather-cep/internal/models"
)

type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperature(city string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Current.TempC, nil
}

func ConvertTemperatures(tempC float64) models.Temperature {
	tempF := tempC*1.8 + 32
	tempK := tempC + 273.15

	return models.Temperature{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}
} 