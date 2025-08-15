package neo4j

import (
	"context"
	"encoding/json"
	"fmt"

	cfg "bit/internal/config"
	"bit/internal/db/neo4j/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/config"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
)

type graphRepository struct {
	driver neo4j.DriverWithContext
}

func NewGraphRepository(ctx context.Context, neo4jCfg *cfg.Neo4jConfig) *graphRepository {
	return &graphRepository{}
}

func (gr *graphRepository) InitGraphRepository(ctx context.Context, cfg *cfg.Neo4jConfig) error {
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
		return err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return err
	}

	gr.driver = driver
	return nil
}

func (gr *graphRepository) CreateBit(ctx context.Context, bit *model.Bit) (string, error) {
	result, err := neo4j.ExecuteQuery(ctx, gr.driver, `
	CREATE (b:Bit {
		AuthorId: $AuthorId,
		Name: $Name,
		Length: $Length,
		Path: $Path,
		Tags: $Tags
	})
	RETURN elementId(b) as id
	`,
		map[string]any{
			"AuthorId": bit.AuthorId,
			"Name":     bit.Name,
			"Length":   bit.Length,
			"Path":     bit.Path,
			"Tags":     bit.Tags,
		},
		neo4j.EagerResultTransformer)
	if err != nil {
		return "", err
	}

	summary := result.Summary
	fmt.Printf("Created %v nodes in %+v.\n",
		summary.Counters().NodesCreated(),
		summary.ResultAvailableAfter())

	id := result.Records[0].Values[0].(string)

	return id, nil
}

func (gr *graphRepository) CreateLinkedBit(ctx context.Context, bit *model.Bit, parentBitId string) (string, error) {
	result, err := neo4j.ExecuteQuery(ctx, gr.driver, `
	MATCH (a:Bit) WHERE elementId(a) = $ParentBitId
	CREATE (a)-[:KNOWS]->(b:Bit {
		AuthorId: $AuthorId,
		Name: $Name,
		Length: $Length,
		Path: $Path,
		Tags: $Tags
	})
	CREATE (b)-[:KNOWS]->(a)
	RETURN elementId(b) as id
	`,
		map[string]any{
			"ParentBitId": parentBitId,
			"AuthorId":    bit.AuthorId,
			"Name":        bit.Name,
			"Length":      bit.Length,
			"Path":        bit.Path,
			"Tags":        bit.Tags,
		},
		neo4j.EagerResultTransformer)
	if err != nil {
		return "", err
	}

	summary := result.Summary
	fmt.Printf("Created %v nodes in %+v.\n",
		summary.Counters().NodesCreated(),
		summary.ResultAvailableAfter())

	if len(result.Records) == 0 || len(result.Records[0].Values) == 0 {
		return "", fmt.Errorf("no id returned from query")
	}

	id, ok := result.Records[0].Values[0].(string)
	if !ok {
		return "", fmt.Errorf("unexpected id type: %T", result.Records[0].Values[0])
	}

	return id, nil
}

func (gr *graphRepository) GetBitById(ctx context.Context, bitId string) (*model.Bit, error) {
	session := gr.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
		FetchSize:  1,
	})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
		MATCH (b:Bit) WHERE elementId(b) = $bitId
		RETURN elementId(b) as Id, b.AuthorId as AuthorId, b.CoAuthorIds as CoAuthorIds,
			b.Name as Name, b.Length as Length, b.Path as Path,
			b.Tags as Tags, b.AditionalTags as AditionalTags
		LIMIT 1
	`,
		map[string]any{
			"bitId": bitId,
		},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return nil, err
	}

	bitBytes, err := json.Marshal(record.AsMap())
	if err != nil {
		return nil, err
	}

	var bit model.Bit
	err = json.Unmarshal(bitBytes, &bit)
	if err != nil {
		return nil, err
	}

	return &bit, nil
}

func (gr *graphRepository) GetBranchByBitId(ctx context.Context, bitId string, depth int) {
	result, err := neo4j.ExecuteQuery(ctx, gr.driver, `
		MATCH (root:Bit {elementId(root): $nodeId})
		MATCH path = (root)-[:KNOWS*]->(child)
		RETURN path
	`,
		map[string]any{
			"nodeId": bitId,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("<database-name>"))
	if err != nil {
		return
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
