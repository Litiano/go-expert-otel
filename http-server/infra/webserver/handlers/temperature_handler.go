package handlers

import (
	"encoding/json"
	"errors"
	"go-expert-otel/http-server/infra/dto"
	http2 "go-expert-otel/http-server/infra/http"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func validateCep(cep string) error {
	match, err := regexp.Match("^[0-9]{8}$", []byte(cep))
	if err != nil {
		return err
	}
	if match {
		return nil
	}
	return errors.New("invalid zipcode")
}
func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	tracer := otel.Tracer("microservice-tracer")
	ctx, initialSpan := tracer.Start(ctx, "Start Request")
	defer initialSpan.End()

	requestId := r.Context().Value(middleware.RequestIDKey).(string)
	var input dto.TemperatureInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	err = validateCep(input.Cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	ctx, getTemperatureSpan := tracer.Start(ctx, "GetTemperature")
	resp, err := http2.RequestWithTimeout(ctx, 10*time.Second, "GET", "http://temperature-service:8090/temperature?cep="+input.Cep, nil, requestId)
	//resp, err := http2.RequestWithTimeout(ctx, 10*time.Second, "GET", "http://127.0.0.1:8090/temperature?cep="+input.Cep, nil, requestId)
	getTemperatureSpan.End()
	if err != nil && resp == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil && resp != nil {
		st, _ := io.ReadAll(resp.Body)
		http.Error(w, string(st), resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	var output dto.TemperatureOutput
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(output)
}
