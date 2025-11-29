package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "messages.txt" // Replace with your file name
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed

	chunkSize := 8
	buffer := make([]byte, chunkSize) // Create a byte slice of size 8

	for {
		// Read up to 8 bytes into the buffer
		bytesRead, err := file.Read(buffer)

		if err != nil {
			// Check for end of file, which is normal termination
			if err == io.EOF {
				break // Exit the loop at the end of the file
			}
			// Handle other potential errors
			log.Fatal(err)
		}

		// Process the bytes read.
		// Use buffer[:bytesRead] to work with only the actual number of bytes read
		// in the current iteration, especially if the last chunk is smaller than 8 bytes.
		fmt.Printf("read: %s\n", string(buffer[:bytesRead]))
	}
}
