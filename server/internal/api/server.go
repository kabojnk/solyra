package api

import (
	"context"
	"net/http"

	"github.com/kabojnk/solyra/server/internal/api/v1"
	"github.com/kabojnk/solyra/server/internal/config"
	"github.com/kabojnk/solyra/server/internal/db"
	"github.com/kabojnk/solyra/server/internal/services/cache"
	"github.com/kabojnk/solyra/server/internal/services/weather"

	"github.com/gin-gonic/gin"
)

// Server represents the API server
type Server struct {
	router        *gin.Engine
	httpServer    *http.Server
	db            *db.PostgresDB
	redisClient   *cache.RedisClient
	weatherClient *weather.Client
	config        *config.Config
}

// NewServer creates a new API server
func NewServer(cfg *config.Config, database *db.PostgresDB, redisClient *cache.RedisClient) *Server {
	router := gin.Default()

	// Create weather client
	weatherClient := weather.NewClient(cfg.Weather.APIKey)

	server := &Server{
		router:        router,
		db:            database,
		redisClient:   redisClient,
		weatherClient: weatherClient,
		config:        cfg,
	}

	// Setup routes
	server.setupRoutes()

	return server
}

// setupRoutes sets up the API routes
func (s *Server) setupRoutes() {
	// Health check endpoint
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API v1 routes
	v1API := v1.NewAPI(s.db, s.redisClient, s.weatherClient, s.config)
	v1Group := s.router.Group("/api/v1")
	{
		v1API.RegisterRoutes(v1Group)
	}
}

// Start starts the API server
func (s *Server) Start() error {
	s.httpServer = &http.Server{
		Addr:    s.config.Server.Address,
		Handler: s.router,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the API server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
