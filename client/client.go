package client

import (
	"fmt"
	"net"
	"network"
	"os"
	"server"
	"sync"
)

type Client struct {
	network.Connection
	id         byte
	wg         *sync.WaitGroup
	authorized bool
}

func NewClient(addr string, port string) *Client {
	c := &Client{
		id: 77,
		wg: &sync.WaitGroup{},
	}

	conn, err := net.Dial("tcp", addr+":"+port)
	if err != nil {
		fmt.Println("Error connecting server")
		os.Exit(1)
	}
	c.Conn = conn

	go c.Reader()
	go c.Writer()

	return c
}

func (c *Client) Run() {
	c.wg.Add(1)
	go c.userInput()

	c.wg.Wait()
}

func (c *Client) userInput() {
	defer c.wg.Done()

	for {
		fmt.Print("Chat:")
		var input string
		fmt.Scanln(&input)

		if len(input) == 0 {
			continue
		}

		p := server.PacketClientChatMessage(c.id, input)
		c.Send(p)
	}
}
