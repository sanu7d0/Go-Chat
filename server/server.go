package server

import (
	"log"
	"net"
	"network"
	"os"
	"sync"
)

type Server struct {
	port string

	packetReceive chan network.Packet
	wg            *sync.WaitGroup
}

func NewServer(port string) *Server {
	return &Server{
		port:          port,
		packetReceive: make(chan network.Packet, 512),
		wg:            &sync.WaitGroup{},
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

		_ = NewChatClient(conn, s.packetReceive)

		// TODO: resolve id from client
	}
}

func (s *Server) handlePackets() {
	defer s.wg.Done()

	for packet := range s.packetReceive {
		s.HandlePacket(network.NewPacketReader(&packet))
	}
}

func (s *Server) broadcast(packet network.Packet) {

}
