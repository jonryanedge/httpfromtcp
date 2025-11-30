package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const address = "127.0.0.1"
const port = "42069"

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:42069")
	if err != nil {
		log.Fatal("failed to resolve udp address")
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("error connecting to udp address")
	}
	defer conn.Close()

	b := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		input, err := b.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s", err)
			break
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Printf("Error writing to connection: %s", err)
			break
		}
	}
}
