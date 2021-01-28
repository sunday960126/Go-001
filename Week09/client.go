package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatalf("connect err:%v", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Printf("read from console  err: %v", err)
			break
		}
		input = strings.TrimSpace(input)
		if input == "quit" {
			log.Printf("client quit")
			break
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Printf("write err: %v", err)
			break
		}
	}

}
