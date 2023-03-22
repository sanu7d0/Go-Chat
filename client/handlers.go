package client

import (
	"fmt"
	"network"
	"server"
)

func (c *Client) HandlePacket(r network.PacketReader) {
	switch r.ReadByte() {
	case server.ClientChatMessage:
		c.handleClientChatMessage(r)
	default:
		fmt.Println("Undefined opcode")
	}
}

func (c *Client) handleClientChatMessage(r network.PacketReader) {
	senderId := r.ReadByte()
	message := r.ReadString()

	fmt.Printf("(%d):%s\n", senderId, message)
}
