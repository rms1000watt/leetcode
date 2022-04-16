package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("ERROR: not enough args provided. ${1} == host ${2} == port")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	fmt.Printf("connecting to: %s %s\n", host, port)

	return
}
