package pgsql

import (
	"auth-service/config"
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type PgStore struct {
	DB *sql.DB
}

func NewPgStore(config config.DBConfig, ctx context.Context) (*PgStore, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name,
	)

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to DB")
		return nil, err
	}

	return &PgStore{
		DB: DB,
	}, nil
}

func (ps *PgStore) Do() error {
	return nil
}
