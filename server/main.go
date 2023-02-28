package main

import (
	"server/chat"
)

func main() {
	chatServer := chat.ChatServer{Port: "9000"}
	chatServer.Start()
}
