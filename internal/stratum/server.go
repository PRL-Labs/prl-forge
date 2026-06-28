package stratum

import (
	"encoding/json"
	"log"
	"net"
	"sync"
)

type Server struct {
	address string

	mu      sync.Mutex
	clients map[net.Conn]bool
}

func New(address string) *Server {
	return &Server{
		address: address,
		clients: make(map[net.Conn]bool),
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

		s.Register(conn)

		go s.handleConnection(conn)
	}
}

func (s *Server) Register(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clients[conn] = true
}

func (s *Server) Unregister(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.clients, conn)
}

func (s *Server) broadcastRaw(msg any) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	data = append(data, '\n')

	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.clients {
		if _, err := conn.Write(data); err != nil {
			log.Printf("client write error: %v", err)
			conn.Close()
			delete(s.clients, conn)
		}
	}
}

func (s *Server) Broadcast(job *Job) {
	if job == nil {
		return
	}

	msg := map[string]any{
		"id":     nil,
		"method": "mining.notify",
		"params": []any{
			job.ID,
			job.PrevHash,
			job.Coinb1,
			job.Coinb2,
			job.Merkle,
			job.Version,
			job.NBits,
			job.NTime,
			job.Clean,
		},
	}

	s.broadcastRaw(msg)

	log.Printf("📤 mining.notify sent jobID=%s", job.ID)
}
