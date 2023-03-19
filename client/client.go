package client

import (
	"net"
	"network"
)

type Client struct {
	network.Connection
	id string
}

func NewClient(addr string, port string) *Client {
	c := Client{}
	conn, err := net.Dial("tcp", addr+port)
	if err != nil {

	}
}
