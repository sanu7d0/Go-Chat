package network

import (
	"log"
	"net"
)

func StartListen(port string, onReceive func([]byte)) {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Println("Error binding port: ", err)
		return
	}
	log.Println("Listening on ", listener.Addr())

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Println("Error accepting connection")
			continue
		}

		go handleConnection(connection, onReceive)
	}
}

func handleConnection(connection net.Conn, onReceive func([]byte)) {
	receiveBuffer := make([]byte, 4096)

	for {
		n, err := connection.Read(receiveBuffer)

		if err != nil {
			log.Println("Error reading from client")
			break
		}

		// TODO: Check n > length
		// TODO: zero copy
		go onReceive(receiveBuffer[:n])
	}
}
