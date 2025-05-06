package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host     string `envconfig:"host"`
	Port     int    `envconfig:"port"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	Name     string `envconfig:"name"`
}

func MustConfigDB() (*DBConfig, error) {
	var dc DBConfig
	err := envconfig.Process("db", &dc)
	if err != nil {
		return nil, err
	}
	return &dc, nil
}
