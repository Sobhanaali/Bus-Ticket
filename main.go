package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Sobhanaali/Bus-Ticket/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// load config

	config , err := util.LoadConfig(".")

	if(err != nil) {
		log.Fatal().Err(err).Msg("Cannot load config !")
	}

	if(config.AppDebug == "true") {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		config.DBConnection,
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase,
	)

	dbpool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	
	err = dbpool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Cannot connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()
}