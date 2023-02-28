package login

import (
	"fmt"
	"server/network"
)

const (
	LoginPacket = iota
	SignUpPacket
)

func HandleLoginPacket(packet []byte) {
	id, n := network.ReadString(packet[:])
	pw, _ := network.ReadString(packet[n:])

	// TODO: Decrypt
	fmt.Printf("Login: %s / %s", &id, &pw)
}

func HandleSignUpPacket(packet []byte) {

}

func WriteLoginPacket(buffer []byte, id string, pw string) {
	n := network.WrtieString(buffer[:], id)
	_ = network.WrtieString(buffer[n:], pw)
}
