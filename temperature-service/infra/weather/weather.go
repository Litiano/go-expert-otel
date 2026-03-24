package weather

import (
	"encoding/json"
	"fmt"
	http2 "go-exper-otel/temperature/infra/http"
	"go-exper-otel/temperature/infra/temperature"
	"net/http"
	"net/url"
	"time"
)

type weatherResponse struct {
	Current struct {
		Temperature float64 `json:"temp_c"`
	} `json:"current"`
}

type TemperatureResponse struct {
	TemperatureC float64 `json:"temp_C"`
	TemperatureF float64 `json:"temp_F"`
	TemperatureK float64 `json:"temp_K"`
	City         string  `json:"city"`
}

func GetTemperature(city string, token string) (*TemperatureResponse, *http2.HttpError) {
	urlBuilder, _ := url.Parse("https://api.weatherapi.com/v1/current.json")
	query := urlBuilder.Query()
	query.Add("q", city)
	query.Add("key", token)
	query.Add("aqi", "no")
	urlBuilder.RawQuery = query.Encode()

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(urlBuilder.String())
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, http2.NewHttpError("Temperature request error", http.StatusBadRequest)
	}
	defer resp.Body.Close()

	temp := weatherResponse{}
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, http2.NewHttpError("invalid json response", http.StatusInternalServerError)
	}

	return &TemperatureResponse{
		TemperatureC: temp.Current.Temperature,
		TemperatureF: temperature.CelsiusToFahrenheit(temp.Current.Temperature),
		TemperatureK: temperature.CelsiusToKelvin(temp.Current.Temperature),
		City:         city,
	}, nil
}
