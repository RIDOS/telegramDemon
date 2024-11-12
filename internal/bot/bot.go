package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBotAPI(token string) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	botAPI.Debug = true
	return &Bot{api: botAPI}, nil
}

func (b *Bot) Start(authorizedUserName string) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("Received a message from UserName: %s", update.Message.From.UserName)

			if update.Message.From.UserName != authorizedUserName {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "У вас нет прав для выполнения этой команды.")
				_, err := b.api.Send(msg)
				if err != nil {
					log.Println(err)
					return
				}
				continue
			}

			b.HandleUpdate(update)
		}
	}
}
