package stratumv2

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
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	log.Printf("🚀 Pearl Stratum V2 listening on %s", s.address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	log.Printf("🔌 %s connected", conn.RemoteAddr())

	session := NewSession(conn)
	defer session.Close()

	for {
		msg, err := session.Read()
		if err != nil {
			log.Printf("Disconnected: %v", err)
			return
		}

		log.Printf("RPC >>> %#v", msg)

		method, ok := msg["method"].(string)
		if !ok {
			log.Println("Invalid RPC: missing method")
			continue
		}

		switch method {

		case "mining.authorize":
			if err := HandleAuthorize(session, msg); err != nil {
				log.Println(err)
				return
			}

		default:
			log.Printf("Unhandled method: %s", method)
		}
	}
}