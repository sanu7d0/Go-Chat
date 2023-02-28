package main

import (
	"log"
	"net"
	"regexp"
)

type LoginError struct {
	Message string
}

var UserId string
var Connection net.Conn

func TryLogin(id string) error {
	if matched, _ := regexp.MatchString("(?i)^[a-z0-9]{4,10}$", id); !matched {
		return &LoginError{Message: "ID must be number & alphabets, length 4 ~ 10"}
	}

	UserId = id
	// Pages.SwitchToPage("Chats")
	go connect()

	return nil
}

func connect() {
	var err error
	Connection, err = net.Dial("tcp", ":9000")
	if err != nil {
		log.Println("Error connecting server: ", err)
		return
	}

	// defer Connection.Close()
	go receive()
}

func receive() {
	receiveBuffer := make([]byte, 4096)
	for {
		n, err := Connection.Read(receiveBuffer)

		if err != nil {
			log.Println("Error receiving: ", err)
			return
		}

		log.Println("From server: ", string(receiveBuffer[:n]))
	}
}

func (e *LoginError) Error() string {
	return e.Message
}
