package config

import "sync"

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
	cfg.Grpc.Port = 50051

	cfg.DB.Host = "localhost"
	cfg.DB.Port = 5432
	cfg.DB.Name = "goodsstore"
	cfg.DB.User = "dncc"
	cfg.DB.Password = "dncc"
}
