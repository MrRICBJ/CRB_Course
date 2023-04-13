package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Token  string `yaml:"tg_token"`
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func MustLoad() Config {
	token := viper.GetString("tg_token")
	if token == "" {
		log.Fatal("Bot token is not specified")
	}

	return Config{
		Token:  token,
	}
}
