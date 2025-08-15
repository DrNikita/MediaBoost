package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Neo4jDbConfig struct {
	Uri      string `envconfig:"uri"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	Name     string `envconfig:"name"`
}

func Neo4jConfig() (*Neo4jDbConfig, error) {
	var dbCfg Neo4jDbConfig

	err := envconfig.Process("neo4j", &dbCfg)
	if err != nil {
		return nil, err
	}

	return &dbCfg, nil
}
