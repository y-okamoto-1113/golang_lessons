package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("ここの処理は走らない。")
	os.Exit(99) // exitステータスは自由に選べる。
}
