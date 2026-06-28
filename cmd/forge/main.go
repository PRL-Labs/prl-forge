package main

import (
	"log"
"os"
	"github.com/techobg/prl-forge/internal/app"
)

func main() {
	f, _ := os.Create("forge.log")
	defer f.Close()
log.SetOutput(f)
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
