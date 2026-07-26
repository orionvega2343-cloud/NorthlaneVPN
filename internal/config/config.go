package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type DB struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"ssl_mode"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `env:"DB_PASS"`
}

type Redis struct {
	Addr string `yaml:"addr"`
}

type Notifier struct {
	ExpiresAt time.Duration `yaml:"expires_at"`
}

type Config struct {
	Db       *DB       `yaml:"db"`
	Redis    *Redis    `yaml:"redis"`
	Notifier *Notifier `yaml:"notifier"`
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}
	err = cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		log.Fatal("Error loading config")
	}
	return &cfg
}
