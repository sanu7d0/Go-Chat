package network

type PacketReader struct {
	pos    int
	packet *Packet
}

func NewPacketReader(p *Packet) PacketReader {
	return PacketReader{pos: 0, packet: p}
}

func (r *PacketReader) ReadByte() byte {
	return r.packet.ReadByte(&r.pos)
}

func (r *PacketReader) ReadBytes(size int) []byte {
	return r.packet.ReadBytes(&r.pos, size)
}

func (r *PacketReader) ReadUInt16() uint16 {
	return r.packet.ReadUInt16(&r.pos)
}

func (r *PacketReader) ReadString() string {
	return r.packet.ReadString(&r.pos)
}
