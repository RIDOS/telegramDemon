package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	BotToken           string
	AuthorizedUserName string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Panic("TELEGRAM_BOT_TOKEN is not set")
	}

	authorizedUserName := os.Getenv("AUTHORIZED_USER_NAME")
	if authorizedUserName == "" {
		log.Panic("AUTHORIZED_USER_NAME is not set or is invalid")
	}

	return &Config{
		BotToken:           botToken,
		AuthorizedUserName: authorizedUserName,
	}
}
