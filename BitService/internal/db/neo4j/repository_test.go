package neo4j

import (
	"bit/config"
	"bit/internal/db/neo4j/model"
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var neo4JCreds = &config.Neo4jDbConfig{
	Uri:      "neo4j://localhost",
	User:     "neo4j",
	Password: "password",
}

// func TestRunTest(t *testing.T) {
// 	ctx := context.Background()

// 	testPassword := "letmein!"

// 	neo4jContainer, err := neo4j.Run(ctx,
// 		"neo4j:5",
// 		neo4j.WithAdminPassword(testPassword),
// 		neo4j.WithLabsPlugin(neo4j.Apoc),
// 		neo4j.WithNeo4jSetting("dbms.tx_log.rotation.size", "42M"),
// 	)
// 	defer func() {
// 		if err := testcontainers.TerminateContainer(neo4jContainer); err != nil {
// 			log.Printf("failed to terminate container: %s", err)
// 		}
// 	}()
// 	if err != nil {
// 		log.Printf("failed to start container: %s", err)
// 		return
// 	}
// }

func TestCreateBit(t *testing.T) {
	ctx := context.Background()

	repository, err := NewGraphRepository(ctx, neo4JCreds)
	assert.NoError(t, err)

	testCases := []struct {
		name string
		bit  *model.Bit
	}{
		{
			name: "Success, CreateBit",
			bit: &model.Bit{
				Id:            uuid.NewString(),
				AuthorId:      1,
				Name:          "Any name",
				Length:        180,
				Path:          "path to the file on server",
				Tags:          []model.Tag{model.Electronic, model.KPop, model.JPop, model.Classical},
				AditionalTags: []string{"Self created tag", "Tag which noone never heard about"},
			},
		},
	}

	for _, tc := range testCases {
		id, err := repository.CreateBit(ctx, tc.bit)
		assert.NoError(t, err)
		fmt.Println("---------------------------------------------", id)
	}
}

func TestCreateLinkedBit(t *testing.T) {
	ctx := context.Background()

	repository, err := NewGraphRepository(ctx, neo4JCreds)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		bit         *model.Bit
		parentBitId string
	}{
		{
			name: "Success, CreateBit",
			bit: &model.Bit{
				Id:            uuid.NewString(),
				AuthorId:      3,
				Name:          "NN",
				Length:        181,
				Path:          "path to the file on server",
				Tags:          []model.Tag{model.Electronic, model.KPop, model.JPop, model.Classical},
				AditionalTags: []string{"Self created tag", "Tag which noone never heard about"},
			},
			parentBitId: "4:fe1157a1-8244-434d-8032-c2a604a878a0:1",
		},
	}

	for _, tc := range testCases {
		id, err := repository.CreateLinkedBit(ctx, tc.bit, tc.parentBitId)
		assert.NoError(t, err)
		fmt.Println("_____________________________________", id)
	}
}

func TestGetBitById(t *testing.T) {
	ctx := context.Background()

	repository, err := NewGraphRepository(ctx, neo4JCreds)
	assert.NoError(t, err)

	testCases := []struct {
		name  string
		bitId string
	}{
		{
			name:  "Success, GetBitById",
			bitId: "4:fe1157a1-8244-434d-8032-c2a604a878a0:0",
		},
	}

	for _, tc := range testCases {
		bit, err := repository.GetBitById(ctx, tc.bitId)
		assert.NoError(t, err)
		fmt.Println("_____________________________________", bit.Tags)
	}
}
