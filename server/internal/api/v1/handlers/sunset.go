package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevinmahoney/etrenank/internal/db"
	"github.com/kevinmahoney/etrenank/internal/models"
	"github.com/kevinmahoney/etrenank/internal/photoquality"
	"github.com/kevinmahoney/etrenank/internal/services/cache"
	"github.com/kevinmahoney/etrenank/internal/services/weather"
	"github.com/kevinmahoney/etrenank/internal/config"
)

// SunsetHandler handles sunset quality endpoints
type SunsetHandler struct {
	db            *db.PostgresDB
	redisClient   *cache.RedisClient
	weatherClient *weather.Client
	config        *config.Config
}

// NewSunsetHandler creates a new sunset handler
func NewSunsetHandler(db *db.PostgresDB, redisClient *cache.RedisClient, weatherClient *weather.Client, config *config.Config) *SunsetHandler {
	return &SunsetHandler{
		db:            db,
		redisClient:   redisClient,
		weatherClient: weatherClient,
		config:        config,
	}
}

// GetSunsetQuality handles the sunset quality endpoint
func (h *SunsetHandler) GetSunsetQuality(c *gin.Context) {
	zipCode := c.Param("zipcode")
	if zipCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Zip code is required",
		})
		return
	}

	ctx := c.Request.Context()

	// Try to get from cache first
	cacheKey := fmt.Sprintf("sunset_quality:%s", zipCode)
	
	// Check if cache should be bypassed (only in debug mode)
	noCache := h.config.Server.Debug && c.Query("nocache") == "true"
	
	if !noCache {
		cachedData, err := h.redisClient.Get(ctx, cacheKey)
		if err == nil {
			// Cache hit
			var sunsetQuality models.SunsetQuality
			if err := json.Unmarshal([]byte(cachedData), &sunsetQuality); err == nil {
				log.Printf("[INFO] Serving sunset quality from cache for zipcode: %s", zipCode)
				c.JSON(http.StatusOK, sunsetQuality)
				return
			}
		}
		log.Printf("[INFO] Cache miss for zipcode: %s", zipCode)
	} else {
		log.Printf("[INFO] Cache bypass requested for zipcode: %s", zipCode)
	}

	// Get weather data from API
	log.Printf("[INFO] Fetching fresh weather data from API for zipcode: %s", zipCode)
	weatherData, astronomyData, err := h.weatherClient.GetWeatherByZipCode(zipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch weather data: %v", err),
		})
		return
	}

	// Calculate sunset quality
	overallQuality, factors, interpretation := photoquality.CalculateSunriseQuality(*weatherData, *astronomyData)

	// Create response
	now := time.Now()
	expiresAt := now.Add(1 * time.Hour)

	sunsetQuality := models.SunsetQuality{
		ZipCode:        zipCode,
		OverallQuality: overallQuality,
		Factors:        factors,
		Interpretation: interpretation,
		WeatherData:    *weatherData,
		AstronomyData:  *astronomyData,
		LastUpdated:    now.Format(time.RFC3339),
		ExpiresAt:      expiresAt.Format(time.RFC3339),
	}

	// Cache the result with 1 hour TTL
	jsonData, err := json.Marshal(sunsetQuality)
	if err == nil {
		h.redisClient.Set(ctx, cacheKey, string(jsonData), 1*time.Hour)
	}

	c.JSON(http.StatusOK, sunsetQuality)
}
