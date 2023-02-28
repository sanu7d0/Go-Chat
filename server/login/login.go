package login

import (
	"server/network"
	"sync"
)

var users sync.Map

func init() {
	// TODO: Connect DB
	// users = make(sync.Map[string]string)
}

func NewListener(port string) {
	network.StartListen(port, onReceive)
}

func onReceive(packet []byte) {
	switch int(packet[0]) {
	case LoginPacket:
		go HandleLoginPacket(packet[1:])
	case SignUpPacket:
		go HandleSignUpPacket(packet[1:])
	default:
	}
}
