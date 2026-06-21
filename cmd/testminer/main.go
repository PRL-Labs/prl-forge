package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

type Notify struct {
	Method string   `json:"method"`
	Params []any    `json:"params"`
}

func main() {

	fmt.Println("⛏️ FIXED STRATUM MINER STARTED")

	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	send(conn, `{"id":1,"method":"mining.subscribe","params":[]}`)
	send(conn, `{"id":2,"method":"mining.authorize","params":["test.worker","x"]}`)

	reader := bufio.NewReader(conn)

	var jobID string
	var extranonce string

	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		fmt.Println("⬅️", line)

		// parse job ID from notify
		var msg map[string]any
		json.Unmarshal([]byte(line), &msg)

		if msg["method"] == "mining.notify" {
			params := msg["params"].([]any)

			jobID = params[0].(string)
			extranonce = "00000000"

			fmt.Println("📦 NEW JOB ID:", jobID)
		}

		// AUTO SUBMIT LOOP (simulate mining)
		if jobID != "" {

			for nonce := 0; nonce < 10000; nonce++ {

				submit := fmt.Sprintf(
					`{"id":3,"method":"mining.submit","params":["test.worker","%s","%s","00000000","%08x"]}`,
					jobID,
					extranonce,
					nonce,
				)

				send(conn, submit)
			}
		}
	}
}

func send(conn net.Conn, msg string) {
	fmt.Fprintln(conn, msg)
	fmt.Println("➡️ sent:", msg)
}