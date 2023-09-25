package fyno

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pascaldekloe/jwt"
)

func (app *Application) authenticate(c *gin.Context) {
	app.Logger.Info().Msg("authenticating user")

	c.Header("Vary", "Authorization")
	authorizationHeader := c.GetHeader("Authorization")
	headerParts := strings.Split(authorizationHeader, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is not valid"})
		return
	}

	token := headerParts[1]
	claims, err := jwt.HMACCheck([]byte(token), []byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userID := claims.Subject

	c.Set("userID", userID)
	c.Next()
}
