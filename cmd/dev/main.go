package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	forgeRunning = false
	proxyRunning = false
	testRunning  = false
)

func clear() {
	fmt.Print("\033[2J\033[H")
}

func banner() {
	clear()

	fmt.Println("===================================================")
	fmt.Println("               PRL FORGE DEV CENTER")
	fmt.Println("===================================================")
	fmt.Println()

	status("Forge", forgeRunning)
	status("Proxy", proxyRunning)
	status("Stratum Test", testRunning)

	fmt.Println()
	fmt.Println("RPC            : UNKNOWN")
	fmt.Println("Connected      : 0")
	fmt.Println("Shares         : 0")
	fmt.Println("Blocks         : 0")

	fmt.Println()
	fmt.Println("===================================================")
	fmt.Println("1. Start Forge")
	fmt.Println("2. Start Proxy")
	fmt.Println("3. Start Stratum Test")
	fmt.Println("4. Forge + Test")
	fmt.Println("5. Build Project")
	fmt.Println("6. Clean Screen")
	fmt.Println("7. Exit")
	fmt.Println("===================================================")
	fmt.Print("> ")
}

func status(name string, running bool) {
	if running {
		fmt.Printf("%-15s : RUNNING\n", name)
	} else {
		fmt.Printf("%-15s : STOPPED\n", name)
	}
}

func run(name string, args ...string) error {

	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Start()
}

func build() {

	cmd := exec.Command("go", "build", "./...")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println()
	fmt.Println("Building...")
	fmt.Println()

	err := cmd.Run()

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println()
		fmt.Println("BUILD SUCCESS")
	}

	fmt.Println()
	fmt.Print("Press ENTER...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		banner()

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":

			err := run("go", "run", "./cmd/forge")

			if err != nil {
				log.Println(err)
			} else {
				forgeRunning = true
			}

		case "2":

			err := run("go", "run", "./cmd/proxy")

			if err != nil {
				log.Println(err)
			} else {
				proxyRunning = true
			}

		case "3":

			err := run("go", "run", "./cmd/stratum-test")

			if err != nil {
				log.Println(err)
			} else {
				testRunning = true
			}

		case "4":

			err := run("go", "run", "./cmd/forge")

			if err == nil {
				forgeRunning = true
			}

			time.Sleep(2 * time.Second)

			err = run("go", "run", "./cmd/stratum-test")

			if err == nil {
				testRunning = true
			}

		case "5":

			build()

		case "6":

			clear()

		case "7":

			return
		}
	}
}
