package fyno

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog"
)

func NewDB(logger zerolog.Logger) (*sql.DB, error) {
	logger.Info().Msg("connecting to database")
	dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection is alive
	if err = db.Ping(); err != nil {
		return nil, err
	}

	logger.Info().Msg("connected to database")
	return db, nil
}
