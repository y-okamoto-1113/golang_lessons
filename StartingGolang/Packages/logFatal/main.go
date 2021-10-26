package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	aaa, err := os.Open("fooooo")
	defer aaa.Close()
	fmt.Println("aaa =>", aaa)
	if err != nil {
		log.Fatal("err =>", err)
	}
}
