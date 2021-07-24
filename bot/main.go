package bot

import (
	"MaximusWarden/MaximusGuard/config"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	Picture      = "picture"
	TurnOnAlarm  = "alarm_on"
	TurnOffAlarm = "alarm_off"
)

var EventCh = make(chan string)

type TGBot struct {
	bot *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func (t *TGBot) New() error {
	cfg := config.GetConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.TGBotToken)
	if err != nil {
		return fmt.Errorf("failed to start bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return fmt.Errorf("failed to get updates: %v", err)
	}
	t.updates = updates
	t.bot = bot
	return nil
}

func (t *TGBot) Run() {
	cfg := config.GetConfig()

	for update := range t.updates {
		if chatID := update.Message.Chat.ID; chatID != cfg.ChatID {
			log.Println(t.bot.Send(tgbotapi.NewMessage(chatID, "Unauthorized!")))
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			c := &Command{
				Bot:    t.bot,
				Update: update,
				cfg:    cfg,
			}
			switch update.Message.Command() {
			case Picture:
				c.takePicture()
			case TurnOnAlarm:
				c.turnOnAlarm()
			case TurnOffAlarm:
				c.turnOffAlarm()
			}
		}
	}
}

func (t *TGBot) ServeChannel() {
	eventHandler := EventHandler{}
	log.Println("serving messages channel")
	go func() {
		for event := range EventCh {
			msg, err := eventHandler.ServeEvent(event)
			if err != nil {
				log.Println(err)
				continue
			}
			_, err = t.bot.Send(msg)
			if err != nil {
				log.Println("error while sending message")
			}
		}
	}()
}
