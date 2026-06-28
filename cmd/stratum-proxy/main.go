package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	ListenAddr = ":3333"
	PoolAddr   = "prl-eu.kryptex.network:7048"
)

var logFile *os.File

func main() {
	var err error

	logFile, err = os.Create("proxy.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	ln, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("====================================")
	log.Println("PRL Stratum Proxy")
	log.Println("Listening :", ListenAddr)
	log.Println("Forwarding:", PoolAddr)
	log.Println("====================================")

	for {
		client, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		log.Println("CLIENT CONNECTED:", client.RemoteAddr())

		go handle(client)
	}
}

func handle(client net.Conn) {
	defer client.Close()

	log.Println("HANDLE START")

	server, err := net.Dial("tcp", PoolAddr)
	if err != nil {
		log.Println("Dial error:", err)
		return
	}

	log.Println("CONNECTED TO POOL:", PoolAddr)

	defer server.Close()

	done := make(chan struct{}, 2)

	go pipe(client, server, "MINER >>>", done)
	go pipe(server, client, "POOL  <<<", done)

	<-done

	log.Println("SESSION CLOSED")
}

func pipe(src net.Conn, dst net.Conn, prefix string, done chan struct{}) {
	reader := bufio.NewReader(src)

	for {
		line, err := reader.ReadBytes('\n')

		if len(line) > 0 {
			msg := fmt.Sprintf("%s %s", prefix, string(line))

			fmt.Print(msg)

			logFile.WriteString(msg)
			logFile.Sync()

			if _, err2 := dst.Write(line); err2 != nil {
				log.Println(prefix, "write error:", err2)
				break
			}
		}

		if err != nil {
			if err != io.EOF {
				log.Println(prefix, "read error:", err)
			}
			break
		}
	}

	done <- struct{}{}
}
