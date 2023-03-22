package server

import (
	"fmt"
	"network"
)

func (s *Server) HandlePacket(r network.PacketReader) {
	switch r.ReadByte() {
	case ClientChatMessage:
		s.handleClientChatMessage(r)
	}
}

func (s *Server) handleClientChatMessage(r network.PacketReader) {
	senderId := r.ReadByte()
	message := r.ReadString()
	fmt.Printf("[ChatMessage] (%d):%s\n", senderId, message)

	p := PacketClientChatMessage(senderId, message)
	s.defaultRoom.Broadcast(p)
}
