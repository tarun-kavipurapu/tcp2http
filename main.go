package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		fmt.Errorf("Error Opening file %v", err)
		return
	}
	defer file.Close()
	buff := make([]byte, 8)
	currentLine := ""
	for {

		n, err := file.Read(buff)
		if err != nil {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			}
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(buff[:n])
		parts := strings.Split(str, "\n")

		for i := 0; i < len(parts)-1; i++ {
			currentLine += parts[i]
			fmt.Printf("read: %s\n", currentLine)
			currentLine = ""
		}
		currentLine += parts[len(parts)-1]

	}

}
