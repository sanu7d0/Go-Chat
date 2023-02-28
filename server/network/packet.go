package network

func ReadString(packet []byte) (string, int) {
	size := int(packet[0])
	return string(packet[1 : 1+size]), size
}

func WrtieString(buffer []byte, str string) int {
	// size := utf8.RuneCountInString(str)
	size := len(str)
	// TODO: Check size > byte.maximum

	buffer[0] = byte(size)
	copy(buffer[1:], str)

	return int(size + 1)
}
