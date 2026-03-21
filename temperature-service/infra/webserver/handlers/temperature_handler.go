package handlers

import (
	"encoding/json"
	"fmt"
	"go-exper-otel/temperature/configs"
	"go-exper-otel/temperature/infra/viacep"
	"go-exper-otel/temperature/infra/weather"
	"net/http"
)

type TemperatureHandler struct {
	Config *configs.Conf
}

func NewTemperatureHandler(config *configs.Conf) *TemperatureHandler {
	return &TemperatureHandler{Config: config}
}

func (t *TemperatureHandler) TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	address, httpError := viacep.GetAddressViaCepApi(r.URL.Query().Get("cep"))
	if httpError != nil {
		fmt.Println(httpError)
		w.WriteHeader(httpError.Code)
		w.Write([]byte(httpError.Message))
		return
	}
	temperature, httpError := weather.GetTemperature(fmt.Sprintf("%s-%s", address.City, address.State), t.Config.WeatherApiKey)
	if httpError != nil {
		fmt.Println(httpError)
		w.WriteHeader(httpError.Code)
		w.Write([]byte(httpError.Message))
		return
	}

	err := json.NewEncoder(w).Encode(temperature)
	if err != nil {
		fmt.Println(httpError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Json encode error"))
		return
	}
}
