package mqtt_client

import (
	"MaximusWarden/MaximusGuard/config"
	"MaximusWarden/MaximusGuard/helpers"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
)

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	cfg := config.GetConfig()
	val := string(msg.Payload())
	log.Println(val)
	if val == "1" {
		log.Println("bot stuff")
		reader := helpers.TakePicture()
		bts, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Println("Error while casting bytes")
			return
		}
		MessageCh <- tgbotapi.NewMessage(cfg.ChatID, "Alert moving detected")
		MessageCh <- tgbotapi.NewPhotoUpload(
			cfg.ChatID,
			tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: bts,
			},
		)
	}
}

var connectionHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected to mqtt")
	if token := client.Subscribe("alarm/value", 0, messageHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func SetupOpts() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(":1883")
	opts.SetClientID("go_alert_client")
	opts.OnConnect = connectionHandler
	return opts
}
