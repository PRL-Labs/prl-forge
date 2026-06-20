package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	fmt.Println("Connected!")

	// -----------------------------
	// Subscribe
	// -----------------------------
	subscribe := `{"id":1,"method":"mining.subscribe","params":[]}` + "\n"

	fmt.Println(">>>", strings.TrimSpace(subscribe))

	if _, err := conn.Write([]byte(subscribe)); err != nil {
		panic(err)
	}

	// -----------------------------
	// Authorize
	// -----------------------------
	authorize := `{"id":2,"method":"mining.authorize","params":["wallet.worker","x"]}` + "\n"

	fmt.Println(">>>", strings.TrimSpace(authorize))

	if _, err := conn.Write([]byte(authorize)); err != nil {
		panic(err)
	}

	submitSent := false

	for {
		reply, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected:", err)
			return
		}

		reply = strings.TrimSpace(reply)

		fmt.Println("<<<", reply)

		// След първото mining.notify изпращаме submit
		if !submitSent && strings.Contains(reply, `"method":"mining.notify"`) {

			submit := `{"id":3,"method":"mining.submit","params":["wallet.worker","1","00000001","68555555","deadbeef"]}` + "\n"

			fmt.Println(">>>", strings.TrimSpace(submit))

			if _, err := conn.Write([]byte(submit)); err != nil {
				panic(err)
			}

			submitSent = true
		}
	}
}