package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleMessage(msg *tgbotapi.Message) {
	response := fmt.Sprintf("Olá, %s! Lembre-se de beber água!", msg.From.UserName)
	b.SendMessage(msg.Chat.ID, response)
}

func (b *Bot) SendMessage(CHATID int64, message string) {
	msg := tgbotapi.NewMessage(CHATID, message)
	if _, err := b.API.Send(msg); err != nil {
		fmt.Printf("Erro ao enviar essa mensagem: %v\n", err)
	}
}
