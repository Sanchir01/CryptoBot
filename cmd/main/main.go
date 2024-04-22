package main

import (
	"os"

	telegramBot "github.com/Sanchir01/CryptoBot/pkg/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	binanceApi "github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	setupEnv()

	client := binanceApi.NewClient(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	//initTGBot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN_BOT"))
	if err != nil {
		logrus.Fatal(err)
	}

	bot.Debug = true

	myBot := telegramBot.NewClientBot(bot, client)

	if err := myBot.Start(); err != nil {
		logrus.Fatal(err)
	}
}

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}
}
