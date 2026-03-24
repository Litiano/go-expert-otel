package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	http2 "go-exper-otel/temperature/infra/http"
	"net/http"
	"regexp"
	"time"
)

type AddressViaCepApi struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

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

func GetAddressViaCepApi(cep string) (*AddressViaCepApi, *http2.HttpError) {
	err := validateCep(cep)
	if err != nil {
		return nil, http2.NewHttpError("invalid zipcode", http.StatusUnprocessableEntity)
	}

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, http2.NewHttpError("error getting address", http.StatusBadRequest)
	}
	defer resp.Body.Close()

	addr := AddressViaCepApi{}
	err = json.NewDecoder(resp.Body).Decode(&addr)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, http2.NewHttpError("invalid json response", http.StatusInternalServerError)
	}

	if addr.Cep == "" {
		return nil, http2.NewHttpError("can not find zipcode", http.StatusNotFound)
	}

	return &addr, nil
}
