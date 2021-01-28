package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const address = ":1234"

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}
	log.Printf("listening at: %s", address)

	s := make(chan os.Signal)
	go func() {
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		<-s

		if err = listen.Close(); err != nil {
			log.Printf("listen close error: %v", err)
		}
		log.Print("listen close")
	}()

	msgChan := make(chan string)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("accept err: %v", err)
		}
		go readMessage(conn, msgChan)
		go writeMessage(conn, msgChan)
	}
}

func readMessage(conn net.Conn, input chan<- string) {
	defer conn.Close()

	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("read message err: %v", err)
			break
		}
		msg := string(buf[:n])
		log.Printf("read message : %v", msg)
		input <- msg
	}
}

func writeMessage(conn net.Conn, output <-chan string) {
	defer conn.Close()

	for {
		msg := <-output
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Printf("write message err: %v", err)
		}
	}
}
