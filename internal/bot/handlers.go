package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"telegramDemon/internal/computer"
)

func (b *Bot) HandleUpdate(update tgbotapi.Update) {
	switch update.Message.Text {
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Команды: \n\n\n"+
			"/shutdown - Выключение системы\n\n"+
			"/sleep - Уход в сон")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
			return
		}
	case "/shutdown":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выключаю компьютер...")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
			return
		}
		computer.ShutdownComputer()
	case "/sleep":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ухожу в сон...")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
			return
		}
		computer.SleepComputer()
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте "+
			"/computer для выключения компьютера.")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
