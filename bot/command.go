package bot

import (
	"MaximusWarden/MaximusGuard/config"
	"MaximusWarden/MaximusGuard/devices"
	"MaximusWarden/MaximusGuard/queue"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
)

type Command struct {
	Bot             *tgbotapi.BotAPI
	Update          tgbotapi.Update
	cfg             *config.Config
}

func (c *Command) takePicture() {
	log.Println("Done")
	camera := devices.Camera{}
	reader := camera.TakePicture()
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
	queue.EventCh <- queue.AlarmON
	c.Bot.Send(tgbotapi.NewMessage(c.cfg.ChatID, "Alarm is ON!"))
}

func (c *Command) turnOffAlarm() {
	queue.EventCh <- queue.AlarmOFF
	c.Bot.Send(tgbotapi.NewMessage(c.cfg.ChatID, "Alarm is OFF!"))
}
