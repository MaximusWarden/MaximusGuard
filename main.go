package main

import (
	"MaximusWarden/MaximusGuard/bot"
	mqtt_client "MaximusWarden/MaximusGuard/mqtt-client"
)

func main() {
	client := mqtt_client.Run()
	bot.RunBot(client)
}
