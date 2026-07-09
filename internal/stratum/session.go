package stratum

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

type Session struct {
	conn net.Conn

	Subscribed bool
	Authorized bool

	Wallet string
	Worker string

	Difficulty float64
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn: conn,
	}
}

func (s *Session) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	log.Printf("SEND >>> %s", string(data))

	data = append(data, '\n')

	_, err = s.conn.Write(data)
	return err
}

func (s *Session) Notify(method string, params any) error {
	msg := protocol.Notification{
		ID:     nil,
		Method: method,
		Params: params,
	}

	return s.Send(msg)
}

func (s *Session) Run() {
	defer func() {

		if p := pool.Current(); p != nil && s.Wallet != "" {
			p.Workers().Remove(s.Wallet + "." + s.Worker)

			log.Printf("🗑 Worker removed: %s.%s", s.Wallet, s.Worker)
		}

		log.Println("========== CONNECTION CLOSED ==========")

		_ = s.conn.Close()
	}()

	log.Println("========== NEW CONNECTION ==========")
	log.Printf("Remote: %s", s.conn.RemoteAddr())
	log.Println("Session started")
	go s.jobLoop()

	scanner := bufio.NewScanner(s.conn)

	// Позволява по-големи JSON пакети
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		log.Printf("RAW >>> %q", line)

		s.dispatch(line)
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Printf("Scanner error: %v", err)
	}

	log.Println("Session finished")
}
