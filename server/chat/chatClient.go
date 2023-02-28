package chat

import (
	"log"
	"net"
)

type ChatClient struct {
	Connection net.Conn
}

func (client ChatClient) BeginReceive() {
	receiveBuffer := make([]byte, 4096)

	for {
		n, err := client.Connection.Read(receiveBuffer)

		if err != nil {
			log.Println("Error reading from client")
			client.disconnect()
			break
		}

		// TODO: Check n > length
		// TODO: zero copy
		go onReceive(receiveBuffer[:n])
	}
}

func (client ChatClient) Send(packet []byte) {
	client.Connection.Write(packet)
}

func (client ChatClient) disconnect() {
	client.Connection.Close()
}

func onReceive(buffer []byte) {
	packetId := int(buffer[0])

	switch packetId {
	case ChatMessagePacket:
		HandleChatMessagePacket(buffer[1:])
	}
}
