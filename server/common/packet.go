package common

func ReadString(buffer []byte) (string, int) {
	size := int(buffer[0])
	return string(buffer[1 : 1+size]), size + 1
}

func WrtieString(buffer []byte, str string) int {
	size := len(str)
	// TODO: Check size > byte.maximum

	buffer[0] = byte(size)
	copy(buffer[1:], str)

	return int(size + 1)
}
