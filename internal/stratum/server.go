package stratum

import (
	
	"log"
	"net"
	"sync"
)

type Server struct {
	address string

	mu      sync.Mutex
clients map[net.Conn]*Session
}

func New(address string) *Server {
	return &Server{
		address: address,
		clients: make(map[net.Conn]*Session),
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

func (s *Server) Register(session *Session) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.clients[session.conn] = session
}

func (s *Server) Unregister(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.clients, conn)
}

func (s *Server) Broadcast(job *Job) {
	if job == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for conn, session := range s.clients {

    if err := SendCurrentJob(session); err != nil {
               log.Printf("NotifyJob ERROR for %s: %v", conn.RemoteAddr(), err)
			conn.Close()
			delete(s.clients, conn)
		}
	}
}

	


