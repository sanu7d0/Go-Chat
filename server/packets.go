package server

import "network"

const (
	ClientChatMessage = byte(iota)
	ClientJoinRoom
)

func PacketClientChatMessage(senderId byte, message string) network.Packet {
	p := network.NewPacket(ClientChatMessage)

	p.WriteByte(senderId)
	p.WriteString(message)

	length := uint16(len(p) - network.HeaderSize)
	p.WriteHeader(length)

	return p
}
