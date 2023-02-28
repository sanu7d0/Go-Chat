package main

import (
	"bufio"
	"fmt"
	"os"
	"server/chat"
)

func StartChat() {
	userId := "abcde"

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Message: ")
		input, _ := reader.ReadString('\n')

		size := len(userId) + len(input) + 1
		packet := make([]byte, size)
		chat.WriteChatMessagePacket(packet, userId, input)

		Connection.Write(packet)
	}
}
