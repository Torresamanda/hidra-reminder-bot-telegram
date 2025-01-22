package main

import (
	"fmt"
	"hidra-bot-reminder/config"
	"hidra-bot-reminder/telegram"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações: %v", err)
	}

	bot, err := telegram.NewBot(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Erro ao iniciar o bot: %v", err)
	}

	fmt.Println("Bot iniciando. Aguardando mensagens...")
	bot.Start()
}
