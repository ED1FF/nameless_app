package configs

import (
	"os"
)

type DbConfig struct {
	Username string
	Password string
	Address  string
	Name     string
}

type Config struct {
	Db DbConfig
}

func New() *Config {
	return &Config{
		Db: DbConfig{
			Username: getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Address:  getEnv("DB_ADDRESS", "localhost:5432"),
			Name:     getEnv("DB_NAME", "nameless_app"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
