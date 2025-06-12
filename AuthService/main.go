package main

import (
	"auth-service/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	_, err := config.MustConfigApp()
	if err != nil {
		log.Fatal().Err(err).Msg("˙†˙App config error˙†˙")
	}
	_, err = config.MustConfigDB()
	if err != nil {
		log.Fatal().Err(err).Msg("˙†˙DB config error˙†˙")
	}

	_, err = config.MustConfigEventbus()
	if err != nil {
		log.Fatal().Err(err).Msg("˙†˙DB config error˙†˙")
	}

	// lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port))
	// if err != nil {

	// }
	// var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials: %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }

	// grpcServer := grpc.NewServer()
	// authServer := api.NewAuthenticator()
	// pb.RegisterAuthServer(grpcServer, authServer)
	// grpcServer.Serve(lis)
}
