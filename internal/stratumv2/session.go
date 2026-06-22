package stratumv2

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"sync"
)

type Session struct {
	conn net.Conn

	reader *bufio.Reader
	writer *bufio.Writer

	mu sync.Mutex
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
}

func (s *Session) Close() error {
	return s.conn.Close()
}

func (s *Session) Send(v any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	b = append(b, '\n')

	log.Printf("SEND >>> %s", string(b))

	if _, err := s.writer.Write(b); err != nil {
		return err
	}

	return s.writer.Flush()
}

func (s *Session) Read() (map[string]any, error) {
	line, err := s.reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	log.Printf("RAW >>> %s", string(line))

	var msg map[string]any

	if err := json.Unmarshal(line, &msg); err != nil {
		return nil, err
	}

	return msg, nil
}
