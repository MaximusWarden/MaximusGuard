package bot

import (
	"MaximusWarden/MaximusGuard/config"
	mqtt_client "MaximusWarden/MaximusGuard/mqtt-client"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	Picture      = "picture"
	TurnOnAlarm  = "alarm_on"
	TurnOffAlarm = "alarm_off"
	Ping         = "ping"
)

func InitBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	cfg := config.GetConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.TGBotToken)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	return bot, updates
}

func RunBot(client mqtt.Client) {
	bot, updates := InitBot()
	cfg := config.GetConfig()
	go func() {
		log.Println("serving messages")
		for msg := range mqtt_client.MessageCh {
			_, err := bot.Send(msg)
			if err != nil {
				log.Println("error while sending message")
			}
		}
	}()

	for update := range updates {
		c := &Command{
			Bot:             bot,
			Update:          update,
			mqttGuardClient: client,
			cfg:             cfg,
		}
		if chatID := update.Message.Chat.ID; chatID != cfg.ChatID {
			log.Println(bot.Send(tgbotapi.NewMessage(chatID, "Unauthorized!")))
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case Picture:
				c.takePicture()
				break
			case TurnOnAlarm:
				c.turnOnAlarm()
				break
			case TurnOffAlarm:
				c.turnOffAlarm()
				break
			case Ping:
				c.ping()
				break
			}
		}
	}
}
