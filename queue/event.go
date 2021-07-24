package queue

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	AlarmON = "alarmON"
	AlarmOFF = "alarmOFF"
)

const (
	alarmTopic = "alarm/switch"
)

type EventHandler struct{
	client mqtt.Client
}

func (e *EventHandler) ServeEvent(event string) error  {
	switch event {
	case AlarmON:
		e.client.Publish(alarmTopic, 0, true, "1")
	case AlarmOFF:
		e.client.Publish(alarmTopic, 0, true, "0")
	default:
		return errors.New("unexpected event")
	}
	return nil
}