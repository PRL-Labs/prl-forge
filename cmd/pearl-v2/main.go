package main

import (
	"log"

	"github.com/techobg/prl-forge/internal/stratumv2"
)

func main() {
	server := stratumv2.New(":3334")

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
