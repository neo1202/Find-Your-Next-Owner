package fyno

import (
	"fyno/server/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Application) NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174", "http://fyno-client-dev.ap-northeast-1.elasticbeanstalk.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/ping", handlers.Ping)

	user := r.Group("/api/users")
	{
		// user.GET("", app.Handlers.Users.GetAllUsers)
		user.GET("/:id", app.Handlers.Users.GetUser)
		user.GET("/name/:name", app.Handlers.Users.GetUserByName)
		user.POST("", app.Handlers.Users.CreateUser)
		user.PUT("/me", app.authenticate, app.Handlers.Users.UpdateUser)
		user.GET("/:id/posts", app.Handlers.Users.GetUserPosts)
	}

	post := r.Group("/api/posts")
	{
		post.GET("", app.Handlers.Posts.GetAllPosts)
		post.GET("/:id", app.Handlers.Posts.GetPost)
		post.POST("", app.authenticate, app.Handlers.Posts.CreatePost)
	}

	r.GET("/api/locations", app.Handlers.Locations.GetAllLocations)
	r.GET("/api/categories", app.Handlers.Categories.GetAllCategories)

	r.GET("/ws/:user_id", app.Handlers.WebSockets.WsConnection)
	r.GET("/ws/connection/:user_id", app.Handlers.WebSockets.IsUserConnected)

	message := r.Group("/api/messages/user_groups")
	{
		message.GET("", app.authenticate, app.Handlers.Messages.GetAllUserGroups)
		message.POST("", app.authenticate, app.Handlers.Messages.CreateUserGroup)
	}

	r.POST("/presigned_url", app.authenticate, app.Handlers.S3.CreatePresignedURL)

	return r
}
