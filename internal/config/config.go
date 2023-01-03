package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBName     string
	Port       string
}

func LoadConfig() *Config {
	env := os.Getenv("ACTIVE_ENV")
	if env == "" {
		env = "dev"
	}

	switch env {
	case "dev":
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	case "prod":
	default:
		log.Fatal("Unknown environment")
	}

	log.Println("Loaded environment variables for", env)

	port := "6004"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return &Config{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       port,
	}
}
