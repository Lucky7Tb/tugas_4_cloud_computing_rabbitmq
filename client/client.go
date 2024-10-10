package client

import (
	"bufio"
	"fmt"
	"strings"

	groupChat "mqtt/example/chat"
)

func NewClient(name string, groupChat *groupChat.Chat, reader *bufio.Reader) {
	for {
		fmt.Printf("%s: ", name)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error read input")
		}
		message := strings.TrimSpace(input)

		if strings.ToLower(message) == "exit" {
			groupChat.SendMessage(fmt.Sprintf("%s: %s", name, "Left the chat"))
			break
		} else {
			groupChat.SendMessage(fmt.Sprintf("%s: %s", name, message))
		}
	}
}
