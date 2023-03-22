package server

import (
	"log"
	"net"
	"network"
	"os"
	"sync"
)

type Server struct {
	port          string
	packetReceive chan network.Packet
	wg            *sync.WaitGroup
	// rooms         map[byte]*Room
	defaultRoom *Room
}

func NewServer(port string) *Server {
	return &Server{
		port:          port,
		packetReceive: make(chan network.Packet, 512),
		wg:            &sync.WaitGroup{},
		// rooms:         map[byte]*Room{},
		defaultRoom: &Room{},
	}
}

func (s *Server) Run() {
	s.wg.Add(1)
	go s.acceptClients()

	s.wg.Add(1)
	go s.handlePackets()

	s.wg.Wait()
}

func (s *Server) acceptClients() {
	defer s.wg.Done()

	listener, err := net.Listen("tcp", "localhost:"+s.port)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Listening on " + listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		c := NewChatClient(conn, s.packetReceive)
		log.Println("Accept client from " + conn.RemoteAddr().String())

		s.defaultRoom.Join(c)

		// TODO: Request authentication to client
	}
}

func (s *Server) handlePackets() {
	defer s.wg.Done()

	for packet := range s.packetReceive {
		go s.HandlePacket(network.NewPacketReader(&packet))
	}
}
