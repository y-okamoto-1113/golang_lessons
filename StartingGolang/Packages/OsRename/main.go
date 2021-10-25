package main

import "os"

func main() {
	os.Rename("test.txt", "test2.txt")
	os.Rename("test2.txt", "test.txt")
}
