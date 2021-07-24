package config

import (
	log "github.com/pion/ion-log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	TGBotToken string `yaml:"tg_bot_token"`
	ChatID int64 `yaml:"chat_id"`
}

var (
	config *Config
)

func GetConfig() *Config {
	if config == nil {
		data, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		config = &Config{}
		err = yaml.Unmarshal(data, config)
		if err != nil {
			log.Errorf("failed to unmarshal: %v", err)
		}
	}
	return config
}