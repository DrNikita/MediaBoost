package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type dbConfig struct {
	Host     string `envconfig:"host"`
	Port     string `envconfig:"port"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	Name     string `envconfig:"name"`
}

func MustConfigDB() (*dbConfig, error) {
	var dc dbConfig
	err := envconfig.Process("db", &dc)
	if err != nil {
		return nil, err
	}
	return &dc, nil
}
