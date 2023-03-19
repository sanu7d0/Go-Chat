package server

import "network"

func (s *Server) HandlePacket(r network.PacketReader) {
	switch r.ReadByte() {
	case ClientChatMessage:
		s.clientChatMessage(r)
	}
}

func (s *Server) clientChatMessage(r network.PacketReader) {
	senderId := r.ReadByte()
	message := r.ReadString()

	p := network.NewPacket(ClientChatMessage)
	p.WriteByte(senderId)
	p.WriteString(message)

	s.broadcast(p)
}
