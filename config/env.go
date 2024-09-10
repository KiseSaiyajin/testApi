package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Host:     getEnv("PUBLIC_HOST", "localhost"),
		Port:     getEnv("PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "1234"),
		SSLMode:  getEnv("SSL_MODE", "disable"),
		DBName:   getEnv("DB_NAME", "postgres"),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
