package bot

import (
	"MaximusWarden/MaximusGuard/config"
	"MaximusWarden/MaximusGuard/devices"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
)

const (
	PictureEv = "picture"
)

type EventHandler struct{}

func (e *EventHandler) ServeEvent(event string) (tgbotapi.Chattable, error) {
	cfg := config.GetConfig()
	switch event {
	case PictureEv:
		camera := devices.Camera{}
		reader := camera.TakePicture()
		bts, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("error while casting bytes: %v", err)
		}
		photo := tgbotapi.NewPhotoUpload(
			cfg.ChatID,
			tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: bts,
			},
		)
		photo.Caption = "Alert moving detected!"
		return photo, nil
	default:
		return nil, errors.New("unexpected event")
	}
}
