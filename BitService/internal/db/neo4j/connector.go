package neo4j

import (
	"bit/config"
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Connect(ctx context.Context, cfg *config.Neo4jDbConfig) (neo4j.DriverWithContext, error) {
	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	driver, err := neo4j.NewDriverWithContext(
		cfg.Uri,
		neo4j.BasicAuth(cfg.User, cfg.Password, ""))
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}

	return driver, nil
}
