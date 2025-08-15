package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Neo4jConfig struct {
	Uri      string `envconfig:"uri"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	Name     string `envconfig:"name"`
}

func Neo4jMustConfig() (*Neo4jConfig, error) {
	var dbCfg Neo4jConfig

	err := envconfig.Process("neo4j", &dbCfg)
	if err != nil {
		return nil, err
	}

	return &dbCfg, nil
}
