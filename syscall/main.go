package main

import (
	"log"
	"syscall"
)

func main() {
	_, err := syscall.Write(0, []byte("aaa"))
	if err != nil {
		log.Println("aaa")
	}
}
