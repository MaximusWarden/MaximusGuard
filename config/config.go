package config

import (
	log "github.com/pion/ion-log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type Config struct {
	TGBotToken string `yaml:"tg_bot_token"`
	ChatID int64 `yaml:"chat_id"`
}

var (
	config *Config
	lock = &sync.Mutex{}
)

func GetConfig() *Config {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
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