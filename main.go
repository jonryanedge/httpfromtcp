package main

import (
	"fmt"
	"log"
	"os"
)

const fileName = "messages.txt" // Replace with your file name

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := getLinesChannel(file)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
