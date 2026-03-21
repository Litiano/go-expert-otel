package handlers

import (
	"encoding/json"
	"fmt"
	"go-exper-otel/temperature/configs"
	"go-exper-otel/temperature/infra/viacep"
	"go-exper-otel/temperature/infra/weather"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type TemperatureHandler struct {
	Config *configs.Conf
}

func NewTemperatureHandler(config *configs.Conf) *TemperatureHandler {
	return &TemperatureHandler{Config: config}
}

func (t *TemperatureHandler) TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	tracer := otel.Tracer("microservice-tracer")
	ctx, initialSpan := tracer.Start(ctx, "Request received")
	defer initialSpan.End()

	ctx, getAddressSpan := tracer.Start(ctx, "Get address from ViaCEP")
	address, httpError := viacep.GetAddressViaCepApi(r.URL.Query().Get("cep"))
	getAddressSpan.End()
	if httpError != nil {
		fmt.Println(httpError)
		w.WriteHeader(httpError.Code)
		w.Write([]byte(httpError.Message))
		return
	}

	ctx, getTemperature := tracer.Start(ctx, "Get temperature from Weather")
	temperature, httpError := weather.GetTemperature(fmt.Sprintf("%s-%s", address.City, address.State), t.Config.WeatherApiKey)
	getTemperature.End()
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
