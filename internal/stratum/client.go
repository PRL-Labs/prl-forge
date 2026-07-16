package stratum

import (
	"log"
	"net"
)

func (s *Server) handleConnection(conn net.Conn) {
	log.Println("========== NEW CONNECTION ==========")
	log.Printf("Remote: %s", conn.RemoteAddr())

	session := NewSession(conn)

s.Register(session)

session.Run()
	
	log.Println("========== CONNECTION CLOSED ==========")
}
