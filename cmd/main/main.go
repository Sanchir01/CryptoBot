package main

import (
	myBinancePKG "github.com/Sanchir01/CryptoBot/pkg/binance"
	telegramBot "github.com/Sanchir01/CryptoBot/pkg/bot"
	"github.com/Sanchir01/CryptoBot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	binanceApi "github.com/adshao/go-binance/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	config, err := config.InitConfig()

	if err != nil {
		logrus.Fatal("Ошибка в инициализации конфиг viper", err)
	}

	//binanceApiCLient
	client := binanceApi.NewClient(config.ApiKeyBinance, config.SecretKeyBinance)
	binanceClient := myBinancePKG.NewBinanceClient(client)
	//initTGBot
	bot, err := tgbotapi.NewBotAPI(config.TokenBot)
	if err != nil {
		logrus.Error("Ошибка при инициализации бота ", err)
	}
	bot.Debug = true

	myBot := telegramBot.NewClientBot(bot, binanceClient, config)

	
	if err = myBot.Start(); err != nil {
		logrus.Error("ошибка при передачи в бота старта", err)
	}

}
