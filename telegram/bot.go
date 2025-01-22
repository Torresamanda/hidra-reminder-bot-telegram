package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{API: api}, nil
}

func (b *Bot) Start() {
	updates := b.API.GetUpdatesChan(tgbotapi.NewUpdate(0))

	for update := range updates {
		if update.Message != nil {
			go b.HandleMessage(update.Message)
		}
	}
}
