package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TokenBot         string
	ApiKeyBinance    string
	SecretKeyBinance string

	Messages Messages
}

type Messages struct {
	Commands
	Errors
}

type Commands struct {
	StartCommand        CommandsComponents `mapstructure:"start"`
	BinanceStartCommand CommandsComponents `mapstructure:"binance_start"`
	UnknownCommand      string             `mapstructure:"unknown_command"`
}

type CommandsComponents struct {
	Text        string `mapstructure:"text"`
	Description string `mapstructure:"description"`
}
type Errors struct {
	BinanceOneValueError string
	OneCryptoError       string
}

func InitConfig() (*Config, error) {
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
	if err := viper.BindEnv("token_bot"); err != nil {
		return err
	}
	if err := viper.BindEnv("secret_key_binance"); err != nil {
		return err
	}
	if err := viper.BindEnv("api_key_binance"); err != nil {
		return err
	}
	cfg.TokenBot = viper.GetString("token_bot")
	cfg.SecretKeyBinance = viper.GetString("secret_key_binance")
	cfg.ApiKeyBinance = viper.GetString("api_key_binance")

	return nil
}
