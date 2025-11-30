package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)

		chunkSize := 8
		var currentLine string

		for {
			buffer := make([]byte, chunkSize) // Create a byte slice of size 8
			bytesRead, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					ch <- currentLine
				}
				if errors.Is(err, io.EOF) {
					break // Exit the loop at the end of the file
				}
				fmt.Printf("error: %s\n", err)
				break
			}
			str := string(buffer[:bytesRead])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				line := currentLine + parts[i]
				ch <- line
				currentLine = ""
			}

			currentLine += parts[len(parts)-1]
		}
	}()

	return ch
}
