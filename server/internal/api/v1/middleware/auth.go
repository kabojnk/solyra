package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinmahoney/etrenank/internal/db"
)

// AuthMiddleware handles authentication
type AuthMiddleware struct {
	db *db.PostgresDB
}

// NewAuthMiddleware creates a new auth middleware
func NewAuthMiddleware(db *db.PostgresDB) *AuthMiddleware {
	return &AuthMiddleware{
		db: db,
	}
}

// Authenticate authenticates requests using client ID and secret
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.GetHeader("X-Client-ID")
		clientSecret := c.GetHeader("X-Client-Secret")

		if clientID == "" || clientSecret == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authentication credentials",
			})
			c.Abort()
			return
		}

		// Get application from database
		app, err := m.db.GetApplicationByClientID(clientID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid client ID",
			})
			c.Abort()
			return
		}

		// Validate client secret
		if app.ClientSecret != clientSecret {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid client secret",
			})
			c.Abort()
			return
		}

		// Set application ID in context
		c.Set("application_id", app.ID)
		c.Next()
	}
}
