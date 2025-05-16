package main

import (
	"mediaboost/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	appConfig, err := config.MustConfigApp()
	if err != nil {
		log.Fatal().Err(err).Msg("˙†˙App config error˙†˙")
	}
	dbConfig, err := config.MustConfigDB()
	if err != nil {
		log.Fatal().Err(err).Msg("˙†˙DB config error˙†˙")
	}

	log.Debug().Msg(appConfig.Host + dbConfig.Host)
}
