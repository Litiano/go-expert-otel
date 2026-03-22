package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Conf struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

var cfg *Conf

func LoadConfig(path string) *Conf {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	err := viper.BindEnv("WEATHER_API_KEY")
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Using defaults.")
		} else {
			panic(err)
		}
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
