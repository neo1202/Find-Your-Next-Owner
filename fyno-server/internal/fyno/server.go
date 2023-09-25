package fyno

import "os"

func (app *Application) Serve() error {
	// Create a new router and register all of your application's routes.
	router := app.NewRouter()
	port := os.Getenv("PORT")
	// Start the HTTP server and listen for requests.
	app.Logger.Info().Msg("starting server on port " + port)
	return router.Run(":" + port)
}
