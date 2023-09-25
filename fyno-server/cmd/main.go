package main

import (
	"fyno/server/internal/fyno"

	_ "github.com/lib/pq"
)

func main() {
	logger, err := fyno.NewLogger()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create logger")
	}

	db, err := fyno.NewDB(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
	}
	defer db.Close()

	redisClient, err := fyno.NewRedis(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to redis")
	}
	defer redisClient.Close()

	app, err := fyno.NewApplication(db, redisClient, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create application")
	}

	if err := app.Serve(); err != nil {
		logger.Fatal().Err(err).Msg("failed to serve")
	}
}
