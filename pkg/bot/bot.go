package telegramBot

import (
	myBinancePKG "github.com/Sanchir01/CryptoBot/pkg/binance"
	"github.com/Sanchir01/CryptoBot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	binance *myBinancePKG.BinanceStruct
	config  *config.Config
}

func NewClientBot(bot *tgbotapi.BotAPI, binance *myBinancePKG.BinanceStruct, config *config.Config) *Bot {
	return &Bot{bot: bot, binance: binance, config: config}
}

func (b *Bot) Start() error {
	logrus.Printf("Authorized on account %s", b.bot.Self.UserName)
	cmdCfg, err := b.handleMenuCommandsInit(b.config)

	if err != nil {
		logrus.Fatal("Ошибка в инициализации бокового меню комманд", err)
	}
	b.bot.Send(cmdCfg)

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
