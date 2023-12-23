package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port         int           `env:"SERVER_PORT,reuired"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
	Debug    bool   `env:"DB_DEBUG,required"`
}

func New() *Config {
	var cfg Config
	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatalf("Failed to decode config: %s", err)
	}
	return &cfg
}

func NewDB() *DBConfig {
	var db DBConfig
	if err := envdecode.StrictDecode(&db); err != nil {
		log.Fatalf("Failed to decode DB config: %s", err)
	}
	return &db
}
