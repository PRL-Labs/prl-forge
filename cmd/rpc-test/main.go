package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/techobg/prl-forge/internal/pearl"
)

func main() {

	client := pearl.New(pearl.RPCConfig{
		Host:     "127.0.0.1",
		Port:     44107, // Смени с RPC порта на Pearl
		User:     "forge",
		Password: "forge123",
	})

	template, err := client.GetBlockTemplate()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}