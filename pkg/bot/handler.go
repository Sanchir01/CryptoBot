package telegramBot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	binanceStart = "bin"
)

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case binanceStart:
		return b.handleBinanceStart(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleBinanceStart(message *tgbotapi.Message) error {
	orders, err := b.binance.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, o := range orders {
		fmt.Println(o)
	}
	res, err := b.binance.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(res)
	msg := tgbotapi.NewMessage(message.Chat.ID, "Это команда для бинанса ")

	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) (err error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Привет ты ввел команду старт")
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "я не знаю такой команды")
	_, err := b.bot.Send(msg)
	return err
}
