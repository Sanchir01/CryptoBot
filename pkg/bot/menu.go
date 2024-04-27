package telegramBot

import (
	"github.com/Sanchir01/CryptoBot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMenuCommandsInit(config *config.Config) (tgbotapi.SetMyCommandsConfig, error) {
	cmdCfg := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     config.Messages.StartCommand.Text,
			Description: config.Messages.StartCommand.Description,
		},
		tgbotapi.BotCommand{
			Command:     config.Messages.BinanceStartCommand.Text,
			Description: config.Messages.BinanceStartCommand.Description,
		},
	)
	return cmdCfg, nil
}
