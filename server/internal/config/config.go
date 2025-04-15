package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Weather  WeatherConfig
}

// ServerConfig holds the server configuration
type ServerConfig struct {
	Address string
}

// DatabaseConfig holds the database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// RedisConfig holds the Redis configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

// WeatherConfig holds the weather API configuration
type WeatherConfig struct {
	APIKey string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	redisPort, err := strconv.Atoi(getEnv("REDIS_PORT", "6379"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_PORT: %v", err)
	}

	return &Config{
		Server: ServerConfig{
			Address: getEnv("SERVER_ADDRESS", ":8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "etrenank"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     redisPort,
			Password: getEnv("REDIS_PASSWORD", ""),
		},
		Weather: WeatherConfig{
			APIKey: getEnv("WEATHER_API_KEY", ""),
		},
	}, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
