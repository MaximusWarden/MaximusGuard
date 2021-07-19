package mqtt_client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var MessageCh = make(chan tgbotapi.Chattable)

func Run() mqtt.Client {
	log.Println("Running mqtt-client mqtt-client-client")
	opts := SetupOpts()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
