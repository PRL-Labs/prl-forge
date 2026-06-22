package main

import (
	"log"

	"github.com/techobg/prl-forge/internal/proxy"
)

func main() {
	p := proxy.New(
		":3334",
		"prl-eu.kryptex.network:7048",
	)

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}