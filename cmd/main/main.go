package main

import (
	telegramBot "github.com/Sanchir01/CryptoBot/pkg/bot"
	"github.com/Sanchir01/CryptoBot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	binanceApi "github.com/adshao/go-binance/v2"
	"github.com/sirupsen/logrus"
)

func main() {

	config, err := config.InitConfig()

	if err != nil {
		logrus.Fatal("Ошибка в инициализации конфиг viper", err)
	}
	logrus.Println(config.BotToken, config.ApiKeyBinance, config.SecretKeyBinance)
	//binanceApiCLient
	client := binanceApi.NewClient(config.ApiKeyBinance, config.SecretKeyBinance)

	//initTGBot
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		logrus.Fatal("Ошибка при инициализации бота", err)
	}

	bot.Debug = true

	myBot := telegramBot.NewClientBot(bot, client, config.Messages)

	if err = myBot.Start(); err != nil {
		logrus.Fatal("ошибка при передачи в бота старта", err)
	}

}
