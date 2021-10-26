package main

import (
	"fmt"
	"os"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("host =>", host)
	fmt.Println("\n==================================================================\n")

	for i, env := range os.Environ() {
		fmt.Println(i, env)
	}
	fmt.Println("\n==================================================================\n")

	env := os.Getenv("HOME")
	fmt.Println(env)
	fmt.Println("\n==================================================================\n")

	os.Setenv("EDITOR", "vim")
	fmt.Println("os.Getenv('EDITOR') =>", os.Getenv("EDITOR"))
	fmt.Println("\n==================================================================\n")

	fmt.Println("os.Getpid() =>", os.Getpid())
}
