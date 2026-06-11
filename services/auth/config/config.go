package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret string
	TTL    time.Duration
}

func NewConfig() *Config {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.name", "auth")
	viper.SetDefault("jwt.secret", "super-secret-key-123")
	viper.SetDefault("jwt.ttl", "24h")

	return &Config{
		Server: ServerConfig{
			Port: viper.GetString("server.port"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			Name:     viper.GetString("database.name"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("jwt.secret"),
			TTL:    viper.GetDuration("jwt.ttl"),
		},
	}
}


