package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramBotToken string
}

func Init() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_TOKEN"),
	}, nil

}
