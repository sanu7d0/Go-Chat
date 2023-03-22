package server

import (
	"log"
	"network"
)

type Room struct {
	id      byte
	clients []*ChatClient
}

func (r *Room) Join(c *ChatClient) {
	r.clients = append(r.clients, c)
	log.Printf("%s joined room %d\n", c.id, r.id)
}

func (r *Room) Leave(c *ChatClient) bool {
	for i, client := range r.clients {
		if c == client {
			r.clients = remove(r.clients, i)
			return true
		}
	}

	return false
}

func (r *Room) Broadcast(p network.Packet) {
	for _, client := range r.clients {
		client.Send(p)
	}
	log.Printf("Broadcast in room %d\n", r.id)
}

func remove(clients []*ChatClient, i int) []*ChatClient {
	l := make([]*ChatClient, 0)
	l = append(l, clients[:i]...)
	return append(l, clients[i+1:]...)
}
