package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

type JWTConfig struct {
	JWTSecret string
}

type Config struct {
	DatabaseConfig *DatabaseConfig
	JWTConfig      *JWTConfig
}

var CFG *Config

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	config := &Config{
		DatabaseConfig: &DatabaseConfig{
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBName:     os.Getenv("DB_NAME"),
		},
		JWTConfig: &JWTConfig{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
	}

	CFG = config

	return config, nil
}
