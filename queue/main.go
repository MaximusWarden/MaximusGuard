package queue

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

const (
	host     = ":1883"
	clientID = "go_mqtt_cliet"
)

var EventCh = make(chan string)

type MqttClient struct {
	client mqtt.Client
}

func (m *MqttClient) New() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(host)
	opts.SetClientID(clientID)
	opts.OnConnect = connectionHandler
	m.client = mqtt.NewClient(opts)
}

func (m *MqttClient) Run() {
	log.Println("running queue queue-client")
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}


func (m *MqttClient) ServeChannel() {
	eventHandler := EventHandler{
		client: m.client,
	}

	go func() {
		for msg := range EventCh {
			err := eventHandler.ServeEvent(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}