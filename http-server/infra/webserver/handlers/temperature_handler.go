package handlers

import (
	"encoding/json"
	"errors"
	"go-expert-otel/http-server/infra/dto"
	http2 "go-expert-otel/http-server/infra/http"
	"net/http"
	"regexp"
	"time"
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
	var input dto.TemperatureInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateCep(input.Cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	resp, err := http2.RequestWithTimeout(10*time.Second, "GET", "http://temperature-service:8090/temperature?cep="+input.Cep, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
