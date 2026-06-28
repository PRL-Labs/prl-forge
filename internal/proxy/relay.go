package proxy

import (
	"bufio"
	"io"
	"log"
	"net"
)

func Relay(src net.Conn, dst net.Conn, prefix string) {
	reader := bufio.NewReader(src)

	for {
		line, err := reader.ReadBytes('\n')
		if len(line) > 0 {
			log.Printf("%s %s", prefix, string(line))

			if _, werr := dst.Write(line); werr != nil {
				log.Printf("%s write error: %v", prefix, werr)
				return
			}
		}

		if err != nil {
			if err != io.EOF {
				log.Printf("%s read error: %v", prefix, err)
			}
			return
		}
	}
}
