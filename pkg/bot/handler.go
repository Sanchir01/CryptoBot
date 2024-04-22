package telegramBot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const commandStart = "start"

func (b *Bot) handleCommandStart(message *tgbotapi.Message) {
	switch message.Command() {
	case commandStart:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Привет ты ввел команду старт")
		b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "я не знаю такой команды")
		b.bot.Send(msg)
	}
}
