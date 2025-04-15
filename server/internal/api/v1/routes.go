package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinmahoney/etrenank/internal/api/v1/handlers"
	"github.com/kevinmahoney/etrenank/internal/api/v1/middleware"
	"github.com/kevinmahoney/etrenank/internal/db"
	"github.com/kevinmahoney/etrenank/internal/services/cache"
	"github.com/kevinmahoney/etrenank/internal/services/weather"
)

// API represents the v1 API
type API struct {
	db            *db.PostgresDB
	redisClient   *cache.RedisClient
	weatherClient *weather.Client
}

// NewAPI creates a new v1 API
func NewAPI(db *db.PostgresDB, redisClient *cache.RedisClient, weatherClient *weather.Client) *API {
	return &API{
		db:            db,
		redisClient:   redisClient,
		weatherClient: weatherClient,
	}
}

// RegisterRoutes registers the v1 API routes
func (a *API) RegisterRoutes(router *gin.RouterGroup) {
	// Create handlers
	sunsetHandler := handlers.NewSunsetHandler(a.db, a.redisClient, a.weatherClient)

	// Create middleware
	authMiddleware := middleware.NewAuthMiddleware(a.db)

	// Public routes
	router.GET("/health", handlers.HealthCheck)

	// Protected routes
	protected := router.Group("/")
	protected.Use(authMiddleware.Authenticate())
	{
		protected.GET("/sunset_quality/:zipcode", sunsetHandler.GetSunsetQuality)
	}
}
