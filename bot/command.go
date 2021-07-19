package bot

import (
	"MaximusWarden/MaximusGuard/config"
	"MaximusWarden/MaximusGuard/helpers"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
)

type Command struct {
	Bot             *tgbotapi.BotAPI
	Update          tgbotapi.Update
	mqttGuardClient mqtt.Client
	cfg             *config.Config
}

func (c *Command) takePicture() {
	log.Println("Done")
	reader := helpers.TakePicture()
	bts, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println("Error while casting bytes")
		return
	}
	_, err = c.Bot.Send(tgbotapi.NewPhotoUpload(
		c.Update.Message.Chat.ID,
		tgbotapi.FileBytes{
			Name:  Picture,
			Bytes: bts,
		},
	),
	)

	if err != nil {
		log.Println("Error while taking picture")
	}
}

func (c *Command) turnOnAlarm() {
	c.mqttGuardClient.Publish("alarm/switch", 0, true, "1")
	c.Bot.Send(tgbotapi.NewMessage(c.cfg.ChatID, "Alarm is ON!"))
}

func (c *Command) turnOffAlarm() {
	c.mqttGuardClient.Publish("alarm/switch", 0, true, "0")
	c.Bot.Send(tgbotapi.NewMessage(c.cfg.ChatID, "Alarm is OFF!"))
}

func (c *Command) ping() {
	c.Bot.Send(tgbotapi.NewMessage(c.cfg.ChatID, "Pong!"))
}
