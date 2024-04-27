package telegramBot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	switch message.Command() {
	case b.config.Messages.StartCommand.Text:
		return b.handleStartCommand(message)
	case b.config.Messages.BinanceStartCommand.Text:
		return b.handleBinanceStart(message)
	case "bitcoin":
		return b.sendOneCryptocurrencyValue(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleBinanceStart(message *tgbotapi.Message) error {
	crypto, err := b.binance.GetAllPricesCryptocurrency()
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.Errors.OneCryptoError)
		_, err = b.bot.Send(msg)
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(crypto[1]))
	_, err = b.bot.Send(msg)

	return nil
}

func (b *Bot) sendOneCryptocurrencyValue(message *tgbotapi.Message) error {
	cryptoValue, err := b.binance.GetOneCryptocurrency("BTCUSDT")
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.Errors.OneCryptoError)
		_, err = b.bot.Send(msg)
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(cryptoValue))
	_, err = b.bot.Send(msg)
	return nil
}
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.Commands.StartCommand.Description)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.UnknownCommand)
	_, err := b.bot.Send(msg)
	return err
}
