package main

import (
	"log"

	"github.com/techobg/prl-forge/internal/app"
)

func main() {
	if err := app.New().Run(); err != nil {
		log.Fatal(err)
	}
}