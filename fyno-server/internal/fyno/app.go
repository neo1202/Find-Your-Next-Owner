package fyno

import (
	"database/sql"
	"fyno/server/internal/handlers"
	"fyno/server/internal/repositories"
	"fyno/server/internal/services"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

// Application is a struct that holds the state of your application.
type Application struct {
	Handlers *handlers.Handlers
	Logger   zerolog.Logger
}

// NewApplication creates a new instance of the Application struct.
func NewApplication(db *sql.DB, redisClient *redis.Client, logger zerolog.Logger) (*Application, error) {
	repositories := repositories.NewRepositories(db)
	services := services.NewServices(repositories)
	handlers := handlers.NewHandlers(services)

	application := &Application{
		Handlers: handlers,
		Logger:   logger,
	}

	return application, nil
}
