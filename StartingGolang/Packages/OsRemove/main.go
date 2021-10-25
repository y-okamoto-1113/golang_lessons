package main

import "os"

func main() {
	os.Create("test.go")
	os.Remove("test.go")
}
