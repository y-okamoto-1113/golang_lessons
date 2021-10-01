package main

import (
	"fmt"
	"os"
)

// go run ./main.go testestestestes のように入力すると引数として受け取ってくれる。
func main() {
	input := os.Args
	fmt.Printf("入力された値: %v \n", input)
	fmt.Printf("入力された値の型: %T \n", input)
}
