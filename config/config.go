package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
}

func Load() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("нет .env файла, читаю переменные окружения напрямую")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, errors.New("переменная TELEGRAM_BOT_TOKEN не задана")
	}
	return &Config{TelegramToken: token}, nil
}
