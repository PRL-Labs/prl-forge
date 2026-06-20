package stratum

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

type Session struct {
	conn net.Conn

	Subscribed bool
	Authorized bool
	Wallet    string
	Worker    string
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
func (s *Session) Notify(method string, params interface{}) error {
	msg := protocol.Notification{
		ID:     nil,
		Method: method,
		Params: params,
	}

	return s.Send(msg)
}
func (s *Session) Run() {
	defer s.conn.Close()

	log.Println("Session started")

	scanner := bufio.NewScanner(s.conn)

	for scanner.Scan() {
		line := scanner.Text()

		log.Printf("RAW >>> %s", line)

		s.dispatch(line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %v", err)
	}

	log.Println("Session finished")
}