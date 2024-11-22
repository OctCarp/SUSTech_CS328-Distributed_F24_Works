package config

import "sync"

type Config struct {
	ServiceName string `yaml:"service_name" env:"SERVICE_NAME" default:"api_service"`

	DbGrpc struct {
		Server string `yaml:"server" env:"DB_GRPC_SERVER" default:"localhost"`
		Port   int    `yaml:"port" env:"DB_GRPC_PORT" default:"50051"`
	}

	LogGrpc struct {
		Server string `yaml:"server" env:"LOG_GRPC_SERVER" default:"localhost"`
		Port   int    `yaml:"port" env:"LOG_GRPC_PORT" default:"50052"`
	}

	Api struct {
		Host string `yaml:"host" env:"API_HOST" default:"localhost"`
		Port int    `yaml:"port" env:"API_PORT" default:"10880"`
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
	cfg.ServiceName = "api_service"

	cfg.DbGrpc.Server = "localhost"
	cfg.DbGrpc.Port = 50051

	cfg.LogGrpc.Server = "localhost"
	cfg.LogGrpc.Port = 50052

	cfg.Api.Host = "localhost"
	cfg.Api.Port = 9997
}
