package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"regexp"
)

type LoginError struct {
	Message string
}

var UserId string

func TryLogin(id string) error {
	if matched, _ := regexp.MatchString("(?i)^[a-z0-9]{4,10}$", id); !matched {
		return &LoginError{Message: "ID must be number & alphabets, length 4 ~ 10"}
	}

	UserId = id
	go connect()

	return &LoginError{Message: "Login..."}
}

func connect() {
	connection, err := net.Dial("tcp", ":9000")
	if err != nil {
		log.Println("Error connecting server: ", err)
		return
	}

	defer connection.Close()
	go receive(connection)

	Pages.SwitchToPage("Chats")

	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')

		if err != nil {
			println("Error input: ", err)
			break
		}

		connection.Write([]byte(line))
	}
}

func receive(connection net.Conn) {
	receiveBuffer := make([]byte, 4096)
	for {
		n, err := connection.Read(receiveBuffer)

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
