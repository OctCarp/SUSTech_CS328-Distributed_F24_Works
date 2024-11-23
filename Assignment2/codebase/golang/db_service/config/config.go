package config

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Grpc struct {
		Port int `yaml:"port" env:"GRPC_PORT" default:"50051"`
	}
	DB struct {
		Host     string `yaml:"host" env:"POSTGRES_HOST" default:"localhost"`
		Port     int    `yaml:"port" env:"POSTGRES_PORT" default:"5432"`
		Name     string `yaml:"dbname" env:"POSTGRES_DB" required:"true" default:"goodsstore"`
		User     string `yaml:"user" env:"POSTGRES_USER" required:"true" default:"dncc"`
		Password string `yaml:"password" env:"DB_PASSWORD" required:"true" default:"dncc"`
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
	// Load GRPC configurations
	cfg.Grpc.Port = getEnvAsInt("GRPC_PORT", 50151)

	// Load DB configurations
	cfg.DB.Host = getEnv("POSTGRES_HOST", "localhost")
	cfg.DB.Port = getEnvAsInt("POSTGRES_PORT", 5432)
	cfg.DB.Name = getEnv("POSTGRES_DB", "goodsstore")
	cfg.DB.User = getEnv("POSTGRES_USER", "dncc")
	cfg.DB.Password = getEnv("POSTGRES_PASSWORD", "dncc")
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
