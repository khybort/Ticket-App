package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    Port     string
    DBHost   string
    DBPort   string
    DBUser   string
    DBPass   string
    DBName   string
    UIAddress string
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    return &Config{
        Port:     getEnv("PORT", "8000"),
        DBHost:   getEnv("DB_HOST", "db"),
        DBPort:   getEnv("DB_PORT", "5432"),
        DBUser:   getEnv("DB_USER", "postgres"),
        DBPass:   getEnv("DB_PASSWORD", "example"),
        DBName:   getEnv("DB_NAME", "ticketing_app"),
        UIAddress: getEnv("UI_ADDRESS", "http://localhost:80"),
    }
}

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}
