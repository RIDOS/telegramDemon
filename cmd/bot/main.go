package main

import (
	"log"
	"telegramDemon/config"
	"telegramDemon/internal/bot"
)

func main() {
	cfg := config.LoadConfig()

	botAPI, err := bot.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Start(cfg.AuthorizedUserName)
}
