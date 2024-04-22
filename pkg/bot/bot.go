package telegramBot

import (
	"github.com/adshao/go-binance/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	binance *binance.Client
}

func NewClientBot(bot *tgbotapi.BotAPI, binance *binance.Client) *Bot {
	return &Bot{bot: bot, binance: binance}
}

func (b *Bot) Start() error {
	logrus.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()

	if err != nil {
		logrus.Fatal("Ошибка в получении обновления тг бота")
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommandStart(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
