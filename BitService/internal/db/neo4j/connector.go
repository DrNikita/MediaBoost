package neo4j

import (
	cfg "bit/config"
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
)

func Neo4jConnect(ctx context.Context, cfg *cfg.Neo4jDbConfig) (neo4j.DriverWithContext, error) {
	useConsoleLogger := func(level log.Level) func(config *config.Config) {
		return func(config *config.Config) {
			config.Log = log.ToConsole(level)
		}
	}

	driver, err := neo4j.NewDriverWithContext(
		cfg.Uri,
		neo4j.BasicAuth(cfg.User, cfg.Password, ""),
		useConsoleLogger(log.DEBUG),
	)
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}

	return driver, nil
}
