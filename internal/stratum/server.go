package stratum

import (
	"log"
	"net"
)

type Server struct {
	address string
}

func New(address string) *Server {
	return &Server{
		address: address,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	log.Printf("⛏️ Stratum listening on %s", s.address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("🔌 Miner connected: %s", conn.RemoteAddr())

		go s.handleConnection(conn)
	}
}