package config

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	ServiceName string `yaml:"service_name" env:"SERVICE_NAME" default:"api_service"`

	DbGrpc struct {
		Server string `yaml:"server" env:"DB_SERVICE_HOST" default:"localhost"`
		Port   int    `yaml:"port" env:"DB_SERVICE_PORT" default:"50051"`
	}

	LogGrpc struct {
		Server string `yaml:"server" env:"LOGGING_SERVICE_HOST" default:"localhost"`
		Port   int    `yaml:"port" env:"LOGGING_SERVICE_PORT" default:"50052"`
	}

	Api struct {
		Host string `yaml:"host" env:"API_HOST" default:"localhost"`
		Port int    `yaml:"port" env:"API_PORT" default:"8080"`
	}
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
		loadConfig(config)
	})
	return config
}

func loadConfig(cfg *Config) {
	cfg.ServiceName = getEnv("SERVICE_NAME", "api_service")

	cfg.DbGrpc.Server = getEnv("DB_SERVICE_HOST", "localhost")
	cfg.DbGrpc.Port = getEnvAsInt("DB_SERVICE_PORT", 50151)

	cfg.LogGrpc.Server = getEnv("LOGGING_SERVICE_HOST", "localhost")
	cfg.LogGrpc.Port = getEnvAsInt("LOGGING_SERVICE_PORT", 50152)

	cfg.Api.Host = getEnv("API_HOST", "localhost")
	cfg.Api.Port = getEnvAsInt("API_PORT", 15001)
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}
