package main

import (
	"flag"
	"log"
	"net"

	"github.com/armon/go-socks5"
)

func main() {
	flag.Parse()

	conf := &socks5.Config{}

	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Error creating SOCKS5 server: %v", err)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:1080")
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go func() {
			if err := server.ServeConn(conn); err != nil {
				log.Printf("Error serving connection: %v", err)
			}
			conn.Close()
		}()
	}
}
