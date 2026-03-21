package dto

type TemperatureInput struct {
	Cep string `json:"cep"`
}

type TemperatureOutput struct {
	TemperatureC float64 `json:"temp_C"`
	TemperatureF float64 `json:"temp_F"`
	TemperatureK float64 `json:"temp_K"`
	City         string  `json:"city"`
}
