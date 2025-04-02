package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func ReadStream(read io.ReadCloser) <-chan string {
	lineChan := make(chan string)
	go func() {
		defer read.Close()
		defer close(lineChan)
		currLine := ""
		for {
			buff := make([]byte, 8)
			n, err := read.Read(buff)
			if err != nil {
				if currLine != "" {
					lineChan <- currLine
				}
				if errors.Is(err, io.EOF) {
					break
				}

				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(buff[:n])

			parts := strings.Split(str, "\n")

			for i := 0; i < len(parts)-1; i++ {
				lineChan <- fmt.Sprintf("%s%s", currLine, parts[i])
				currLine = ""
			}

			currLine += parts[len(parts)-1]

		}

	}()
	return lineChan
}
