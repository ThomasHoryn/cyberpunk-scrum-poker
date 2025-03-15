package config

import (
	"os"
	"strings"
	"time"
	"strconv"
)

type Config struct {
	ServerPort    string
	MongoURI      string
	Environment   string
	APIKey        string
	AllowedHosts  []string
	HMACSecret    string
	RateLimit     int
	RateLimitWindow time.Duration
}

func Load() *Config {
	return &Config{
		ServerPort:    getEnv("PORT", "8054"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		APIKey:        getEnv("API_KEY", "cyberpunk_dev_2025"),
		AllowedHosts:  strings.Split(getEnv("ALLOWED_HOSTS", "http://localhost:3000"), ","),
		MongoURI:    getEnv("MONGO_URI", "mongodb://root:cyberpunkscrum_password@localhost:27017/cyberpunkscrum?authSource=admin"),
		HMACSecret:    getEnv("HMAC_SECRET", "neonlights2077"),
		RateLimit:     getEnvInt("RATE_LIMIT", 100),
		RateLimitWindow: getEnvDuration("RATE_LIMIT_WINDOW", time.Minute),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	duration, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}
	return duration
}
