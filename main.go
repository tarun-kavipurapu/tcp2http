package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func readStream(read io.ReadCloser, wg *sync.WaitGroup) <-chan string {
	Line := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(Line)
		buff := make([]byte, 8)

		for {
			n, err := read.Read(buff)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
			}
			str := string(buff[:n])
			parts := strings.Split(str, "\n")
		}

	}()
	return Line
}

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		fmt.Printf("Error:", err)
	}
	var wg sync.WaitGroup

	_ = <-readStream(file, &wg)
	// for  := range stream {
	// 	// fmt.Printf(line)
	// }

	wg.Wait()

}
