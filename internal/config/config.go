package config

import (
	"os"
)

type Config struct {
	ServerPort  string
	MongoURI    string
	Environment string
}

func Load() *Config {
	return &Config{
		ServerPort:  getEnv("PORT", "8054"),
		MongoURI:    getEnv("MONGO_URI", "mongodb://root:cyberpunkscrum_password@localhost:27017/cyberpunkscrum?authSource=admin"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
