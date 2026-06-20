package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	fmt.Println("Connected!")

	// --------------------------------------------------
	// mining.subscribe
	// --------------------------------------------------

	subscribe := `{"id":1,"method":"mining.subscribe","params":[]}` + "\n"

	fmt.Println(">>>", subscribe)

	if _, err := conn.Write([]byte(subscribe)); err != nil {
		panic(err)
	}

	reply, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("<<<", reply)

	// --------------------------------------------------
	// mining.authorize
	// --------------------------------------------------

	authorize := `{"id":2,"method":"mining.authorize","params":["wallet.worker","x"]}` + "\n"

	fmt.Println(">>>", authorize)

	if _, err := conn.Write([]byte(authorize)); err != nil {
		panic(err)
	}

	reply, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("<<<", reply)
}