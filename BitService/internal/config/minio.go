package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type MinioConfig struct {
	Host         string `envconfig:"host"`
	Port         string `envconfig:"port"`
	RootUser     string `envconfig:"root_user"`
	RootPassword string `envconfig:"root_password"`
	UseSSL       bool   `envconfig:"use_ssl"`
	BucketName   string `envconfig:"bucket_name"`
}

func MinioMustConfig() (*MinioConfig, error) {
	var dbCfg MinioConfig

	err := envconfig.Process("minio", &dbCfg)
	if err != nil {
		return nil, err
	}

	return &dbCfg, nil
}
