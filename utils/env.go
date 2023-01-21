package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerName string
	ApiPort    string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	SigningKey []byte
)

func LoadEnv() {
	godotenv.Load(".env")
	ApiPort = getEnv("PORT", "8000")

	ServerName = getEnv("SERVER_NAME", "localhost")
	DbUser = getEnv("DB_USER", "user")
	DbPassword = getEnv("DB_PASSWORD", "password")
	DbName = getEnv("DB_NAME", "db")
	DbPort = getEnv("DB_PORT", "5432")

	secretKey := getEnv("SECRET_KEY", "secret")
	SigningKey = []byte(secretKey)
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
