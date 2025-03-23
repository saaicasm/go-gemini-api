package config

import (
	"os"
)

// add functions to handle bool , slice and integer types.

type EnvConfig struct {
	Username  string
	GeminiAPI string
}

type Config struct {
	Env EnvConfig
}

func New() *Config {
	return &Config{
		Env: EnvConfig{
			Username:  getEnv("GITHUB_USERNAME", ""),
			GeminiAPI: getEnv("API_KEY", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
