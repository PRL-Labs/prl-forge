package proxy

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

type Proxy struct {
	listen string
	target string
}

func New(listen, target string) *Proxy {
	return &Proxy{
		listen: listen,
		target: target,
	}
}

func (p *Proxy) Run() error {
	ln, err := net.Listen("tcp", p.listen)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		client, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go p.handle(client)
	}
}

func (p *Proxy) handle(client net.Conn) {
	defer client.Close()

	server, err := net.Dial("tcp", p.target)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	log.Println("Miner connected")

	clientLog, _ := os.Create("client.log")
	serverLog, _ := os.Create("server.log")

	defer clientLog.Close()
	defer serverLog.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		relay(client, server, clientLog, "CLIENT")
	}()

	go func() {
		defer wg.Done()
		relay(server, client, serverLog, "SERVER")
	}()

	wg.Wait()

	log.Println("Connection closed")
}

func relay(src net.Conn, dst net.Conn, logFile *os.File, prefix string) {
	buf := make([]byte, 65536)

	for {
		n, err := src.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			return
		}

		data := buf[:n]

		log.Printf("[%s]\n%s\n", prefix, string(data))

		logFile.WriteString("\n=============================\n")
		logFile.WriteString(prefix)
		logFile.WriteString("\n=============================\n")
		logFile.Write(data)
		logFile.WriteString("\n")

		_, err = dst.Write(data)
		if err != nil {
			return
		}
	}

}
	func (p *Proxy) Start() error {
	return p.Run()

}