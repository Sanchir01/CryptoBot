package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BotToken         string
	ApiKeyBinance    string
	SecretKeyBinance string

	Messages Messages
}

type Messages struct {
	Commands Commands
	Errors   Errors
}

type Commands struct {
	Start          string `mapstructure:"start"`
	BinanceStart   string `mapstructure:"binance_start"`
	UnknownCommand string `mapstructure:"unknown_command"`
}
type Errors struct {
	BinanceOneValueError string
}

func InitConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.commands", &cfg.Messages.Commands); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("bot_token"); err != nil {
		return err
	}
	if err := viper.BindEnv("secret_key_binance"); err != nil {
		return err
	}
	if err := viper.BindEnv("api_key_binance"); err != nil {
		return err
	}
	cfg.ApiKeyBinance = viper.GetString("api_key_binance")
	cfg.SecretKeyBinance = viper.GetString("secret_key_binance")
	cfg.BotToken = viper.GetString("token_bot")

	return nil
}
