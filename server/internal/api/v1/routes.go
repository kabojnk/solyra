package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kabojnk/solyra/server/internal/api/v1/handlers"
	"github.com/kabojnk/solyra/server/internal/api/v1/middleware"
	"github.com/kabojnk/solyra/server/internal/config"
	"github.com/kabojnk/solyra/server/internal/db"
	"github.com/kabojnk/solyra/server/internal/services/cache"
	"github.com/kabojnk/solyra/server/internal/services/weather"
)

// API represents the v1 API
type API struct {
	db            *db.PostgresDB
	redisClient   *cache.RedisClient
	weatherClient *weather.Client
	config        *config.Config
}

// NewAPI creates a new v1 API
func NewAPI(db *db.PostgresDB, redisClient *cache.RedisClient, weatherClient *weather.Client, config *config.Config) *API {
	return &API{
		db:            db,
		redisClient:   redisClient,
		weatherClient: weatherClient,
		config:        config,
	}
}

// RegisterRoutes registers the v1 API routes
func (a *API) RegisterRoutes(router *gin.RouterGroup) {
	// Create handlers
	sunsetHandler := handlers.NewSunsetHandler(a.db, a.redisClient, a.weatherClient, a.config)

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
