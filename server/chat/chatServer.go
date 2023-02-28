package chat

import (
	"fmt"
	"log"
	"net"
)

type ChatServer struct {
	Port string
}

func (server ChatServer) Start() {
	listener, err := net.Listen("tcp", ":"+server.Port)

	if err != nil {
		log.Println("Error binding port: ", err)
		return
	}
	log.Println("Chat server listening on ", listener.Addr())

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Println("Error accepting connection")
			continue
		}

		fmt.Println("Connected with ", connection.RemoteAddr())
		client := ChatClient{Connection: connection}
		go client.BeginReceive()
	}
}
