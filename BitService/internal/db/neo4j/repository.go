package neo4j

import (
	"context"
	"fmt"
	"log"

	"bit/config"
	"bit/internal/db/neo4j/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GraphRepository struct{}

func NewGraphRepository() *GraphRepository {
	return &GraphRepository{}
}

func (gr *GraphRepository) CreateBit(ctx context.Context, driver neo4j.DriverWithContext, bit model.Bit, cfg config.Neo4jDbConfig) error {
	result, err := neo4j.ExecuteQuery(ctx, driver, `
	CREATE (a:Bit {
		Id: $Id,
		AuthorId: $AuthorId,
		Name: $Name,
		Length: $Length,
		Path: $Path,
		Tags: $Tags
	})`,
		map[string]any{
			"Id":       bit.Id,
			"AuthorId": bit.AuthorId,
			"Name":     bit.Name,
			"Length":   bit.Length,
			"Path":     bit.Path,
			"Tags":     bit.Tags,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("<database-name>"))
	if err != nil {
		return err
	}

	summary := result.Summary
	fmt.Printf("Created %v nodes in %+v.\n",
		summary.Counters().NodesCreated(),
		summary.ResultAvailableAfter())

	return nil
}

func (gr *GraphRepository) CreateLinkedBit(ctx context.Context, driver neo4j.DriverWithContext, bit model.Bit, parentBitId string, cfg config.Neo4jDbConfig) error {
	result, err := neo4j.ExecuteQuery(ctx, driver, `
	MATCH (a:Bit) WHERE a.Id = $parentBitId
	CREATE (a)-[:KNOWS]->(b:Bit {
		Id: $Id,
		AuthorId: $AuthorId,
		Name: $Name,
		Length: $Length,
		Path: $Path,
		Tags: $Tags
	})
	CREATE (b)-[:KNOWS]->(a)
	`,
		map[string]any{
			"parentBitId": parentBitId,
			"Id":          bit.Id,
			"AuthorId":    bit.AuthorId,
			"Name":        bit.Name,
			"Length":      bit.Length,
			"Path":        bit.Path,
			"Tags":        bit.Tags,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("<database-name>"))
	if err != nil {
		return err
	}

	summary := result.Summary
	fmt.Printf("Created %v nodes in %+v.\n",
		summary.Counters().NodesCreated(),
		summary.ResultAvailableAfter())

	return nil
}

func (gr *GraphRepository) GetBranchByNodeId(ctx context.Context, driver neo4j.DriverWithContext, nodeId string, cfg config.Neo4jDbConfig) {
	result, err := neo4j.ExecuteQuery(ctx, driver, `
	MATCH (root:Bit {Id: $nodeId})
	MATCH path = (root)-[:KNOWS*]->(child)
	RETURN path
	`,
		map[string]any{
			"nodeId": nodeId,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("<database-name>"))
	if err != nil {
		log.Fatal(err)
	}

	// for _, f := range result.Records {
	// f.AsMap()[""]
	// }

	summary := result.Summary
	fmt.Printf("Created %v nodes in %+v.\n",
		summary.Counters().NodesCreated(),
		summary.ResultAvailableAfter())

	// return nil
}
