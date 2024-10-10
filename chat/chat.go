package chat

import (
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Chat struct {
	mqtt *MQTT.Client
}

func (c *Chat) InitChat(mqtt *MQTT.Client) {
	c.mqtt = mqtt

	if token := (*c.mqtt).Subscribe("group_chat", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println("Cannot join the chat")
		os.Exit(1)
	}
}

func (c *Chat) SendMessage(message string) {
	if token := (*c.mqtt).Publish("group_chat", 0, false, message); token.Wait() && token.Error() != nil {
		fmt.Println("Failed send message")
	}
}
