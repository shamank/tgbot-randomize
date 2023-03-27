package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shamank/tgbot-randomize/app/internal/config"
	bot "github.com/shamank/tgbot-randomize/app/internal/telegram"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		logrus.Error(err)
	}

	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		logrus.Fatal(err)
	}

	botAPI.Debug = true

	bot := bot.NewBot(botAPI)

	go func() {
		if err := bot.StartUp(); err != nil {
			logrus.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
