package main

import (
	"bit/internal/adapter/minio"
	"bit/internal/adapter/neo4j"
	"bit/internal/config"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	neo4jCfg, err := config.Neo4jMustConfig()
	if err != nil {
		log.Fatal(err)
	}
	minioCfg, err := config.MinioMustConfig()
	if err != nil {
		log.Fatal(err)
	}

	neo4j := neo4j.NewGraphRepository(neo4jCfg)
	err = neo4j.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	minio := minio.NewMinioStorage(*minioCfg)
	err = minio.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// bitService := domain.NewBitService(neo4j, minio)
	// bitService.DoSmthng()
}
