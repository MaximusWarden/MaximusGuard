package main

import (
	"MaximusWarden/MaximusGuard/bot"
	"MaximusWarden/MaximusGuard/queue"
	"log"
)

func main() {
	queueClient := queue.MqttClient{}
	botClient := bot.TGBot{}

	queueClient.New()
	err := botClient.New()
	if err != nil {
		log.Fatal(err)
	}



	botClient.ServeChannel()
	queueClient.ServeChannel()


	queueClient.Run()
	botClient.Run()
}
