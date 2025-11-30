package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("could not accept a connection")
		}
		fmt.Println("connection accepted")
		go func(c net.Conn) {
			defer c.Close()
			lines := getLinesChannel(c)
			for line := range lines {
				fmt.Printf("read: %s\n", line)
			}
			fmt.Println("connection closed")
		}(conn)
	}
}
