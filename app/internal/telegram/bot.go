package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shamank/tgbot-randomize/app/internal/commands"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	botAPI   *tgbotapi.BotAPI
	commands *commands.Commands
}

func NewBot(botAPI *tgbotapi.BotAPI, commands *commands.Commands) *Bot {
	return &Bot{
		botAPI:   botAPI,
		commands: commands,
	}
}

func (b *Bot) StartUp() error {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := b.botAPI.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := b.botAPI.Send(msg); err != nil {
			logrus.Error(err)
			return nil
		}
	}

	return nil
}
