package main

import (
	"fmt"
	"syscall"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	stat := &syscall.Stat_t{}
	if err := syscall.Stat("main.go", stat); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("stat.Blksize:", stat.Blksize)
	fmt.Println("stat.Blocks:", stat.Blocks)
	fmt.Println("stat.Ino:", stat.Ino)
	fmt.Println("stat.Size:", stat.Size)
	fmt.Println("#####################################################")

	spew.Dump(stat)
}
