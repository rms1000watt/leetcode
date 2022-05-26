package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// This just means end of stdin
				break
			}
			fmt.Println("ERROR:", err)
			return
		}

		splitString := strings.Split(line, "\t")

		if len(splitString) < 2 {
			continue
		}

		if splitString[0] == "free" {
			continue
		}

		active := splitString[1]

		fmt.Println("Active:", active)
	}
}
