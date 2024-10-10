package main

import (
	"bufio"
	"fmt"
	groupChat "mqtt/example/chat"
	"mqtt/example/client"
	"os"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var clientName string
	fmt.Print("Enter your name: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	clientName = strings.TrimSpace(input)

	mqttOptions := MQTT.NewClientOptions()
	mqttOptions.AddBroker("mqtt://localhost:1883")
	mqttOptions.SetUsername("lucky")
	mqttOptions.SetPassword("lucky101")
	mqttOptions.SetDefaultPublishHandler(func(client MQTT.Client, message MQTT.Message) {
		chat := string(message.Payload())
		if strings.Split(chat, ": ")[0] != clientName {
			fmt.Print("\r\033[K")
			fmt.Printf("%s\n", chat)
			fmt.Printf("%s: ", clientName)
		}
	})

	mqtt := MQTT.NewClient(mqttOptions)
	if token := mqtt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	groupChat := groupChat.Chat{}
	groupChat.InitChat(&mqtt)
	groupChat.SendMessage(fmt.Sprintf("%s: %s", clientName, "Joined the chat"))
	client.NewClient(clientName, &groupChat, reader)
}
