package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	envVars *Environments
)

type Environments struct {
	APIPort         string `mapstructure:"API_PORT"`
	Environment     string `mapstructure:"ENVIRONMENT"`
	FRBaseURL       string `mapstructure:"FR_BASE_URL"`
	FRToken         string `mapstructure:"FR_TOKEN"`
	FRPlataformCode string `mapstructure:"FR_PLATAFORM_CODE"`
	FRRegisteredNum string `mapstructure:"FR_REGISTERED_NUMBER"`
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.SetDefault("API_PORT", "9000")
	viper.SetDefault("ENVIRONMENT", "local")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, using environment variables")
	}

	if err := viper.Unmarshal(&envVars); err != nil {
		fmt.Println("Error unmarshalling config file")
	}

	return envVars
}

func GetEnvVars() *Environments {
	return envVars
}
