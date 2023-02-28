package chat

import (
	"fmt"
	"server/common"
)

const (
	ChatMessagePacket = iota
)

func HandleChatMessagePacket(packet []byte) {
	offset := 0

	id, n := common.ReadString(packet[offset:])
	offset += n

	message, n := common.ReadString(packet[offset:])
	offset += n

	fmt.Printf("From client %s: %s\n", id, message)
}

func WriteChatMessagePacket(buffer []byte, id string, message string) {
	buffer[0] = byte(ChatMessagePacket)
	n := 1
	n += common.WrtieString(buffer[n:], id)
	n += common.WrtieString(buffer[n:], message)
}
