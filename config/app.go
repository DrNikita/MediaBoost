package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type appConfig struct {
	Host string `envconfig:"host"`
	Port string `envconfig:"port"`
}

func MustConfigApp() (*appConfig, error) {
	var ac appConfig
	err := envconfig.Process("app", &ac)
	if err != nil {
		return nil, err
	}
	return &ac, nil
}
