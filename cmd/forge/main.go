package main

import (
   "io"
  	"log"
"os"
	"github.com/techobg/prl-forge/internal/app"
)





func main() {


	f, _ := os.Create("forge.log")
	defer f.Close()
  
mw := io.MultiWriter(os.Stdout, f)
log.SetOutput(mw)
log.Println("========== BUILD 001 ==========")

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
