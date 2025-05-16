package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Host string `envconfig:"host"`
	Port int    `envconfig:"port"`
}

func MustConfigApp() (*AppConfig, error) {
	var ac AppConfig
	err := envconfig.Process("app", &ac)
	if err != nil {
		return nil, err
	}
	return &ac, nil
}
