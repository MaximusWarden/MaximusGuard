package queue

import (
	"MaximusWarden/MaximusGuard/bot"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

const (
	moveSignal = "1"

)

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	val := string(msg.Payload())
	log.Println(val)
	if val == moveSignal {
		bot.EventCh <- bot.PictureEv
	}
}

var connectionHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected to mqtt")
	if token := client.Subscribe("alarm/value", 0, messageHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
