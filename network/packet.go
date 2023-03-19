package network

type Packet []byte

const (
	HeaderSize = 2
)

func NewPacket(op byte) Packet {
	p := Packet{}
	p.WriteUInt16(0)
	p.WriteByte(op)

	return p
}

func (p *Packet) Size() int {
	return int(len(*p))
}

func (p *Packet) ReadByte(pos *int) byte {
	v := byte((*p)[*pos])
	*pos++
	return v
}

func (p *Packet) ReadBytes(pos *int, length int) []byte {
	bytes := []byte((*p)[*pos : *pos+length])
	*pos += length
	return bytes
}

func (p *Packet) ReadUInt16(pos *int) uint16 {
	return uint16(p.ReadByte(pos)) | (uint16(p.ReadByte(pos)) << 8)
}

func (p *Packet) ReadString(pos *int) string {
	length := int(p.ReadUInt16(pos))
	return string(p.ReadBytes(pos, length))
}

func (p *Packet) WriteByte(value byte) {
	*p = append(*p, value)
}

func (p *Packet) WriteBytes(bytes []byte) {
	*p = append(*p, bytes...)
}

func (p *Packet) WriteUInt16(value uint16) {
	*p = append(*p, byte(value), byte(value>>8))
}

func (p *Packet) WriteString(str string) {
	p.WriteUInt16(uint16(len(str)))
	p.WriteBytes([]byte(str))
}
