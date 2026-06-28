package main

import (
	"log"

	"github.com/techobg/prl-forge/internal/proxy"
)

func main() {
	p := proxy.New(
		":3333",
		"pearl-eu2.luckypool.io:3360",
	)

	log.Println("PRL ProxyDump listening on :3333")
	log.Println("Forwarding to pearl-eu2.luckypool.io:3360")

	if err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
