package network

import "net"

type Connection struct {
	net.Conn
	PacketSend    chan Packet
	PacketReceive chan Packet
	closed        bool
}

func (connection *Connection) Reader() {
	header := true
	readSize := HeaderSize

	for {
		buffer := make([]byte, readSize)

		if _, err := connection.Conn.Read(buffer); err != nil {
			break
		}

		if header {
			readSize = int(buffer[0]) | (int(buffer[1]) << 8)
		} else {
			readSize = HeaderSize
			connection.PacketReceive <- buffer
		}

		header = !header
	}
}

func (connection *Connection) Writer() {
	for {
		packet, ok := <-connection.PacketSend
		if !ok {
			return
		}

		connection.Conn.Write(packet)
	}
}

func (connection *Connection) Send(packet Packet) {
	if connection.closed {
		return
	}

	connection.PacketSend <- packet
}

func (connection *Connection) Disconnect() {
	connection.closed = true
	close(connection.PacketSend)
}
