package server

import (
	"net"
	"network"
)

type ChatClient struct {
	id string
	network.Connection
}

func NewChatClient(conn net.Conn, packetReceive chan network.Packet) *ChatClient {
	c := &ChatClient{}
	c.id = "undefined"
	c.Conn = conn
	c.PacketSend = make(chan network.Packet)
	c.PacketReceive = packetReceive

	go c.Reader()
	go c.Writer()

	return c
}
