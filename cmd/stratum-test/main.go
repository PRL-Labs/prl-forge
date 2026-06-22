package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Request struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

func send(conn net.Conn, v any) {
	data, _ := json.Marshal(v)
	data = append(data, '\n')

	log.Printf("SEND >>> %s", string(data))

	_, err := conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("✅ Connected to Forge")

	send(conn, Request{
		ID:     1,
		Method: "mining.subscribe",
		Params: []any{},
	})

	send(conn, Request{
		ID:     2,
		Method: "mining.authorize",
		Params: map[string]any{
			"wallet": "solo:testwallet.TEST",
			"agent":  "PRL-TestClient/1.0",
			"type":   "v2",
		},
	})

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Println("RECV <<<", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}