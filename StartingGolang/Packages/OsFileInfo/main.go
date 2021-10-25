package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("f =>", f)
	fmt.Println("\n==================================================================\n")

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fi =>", fi)
	fmt.Println("fi.Name() =>", fi.Name())
	fmt.Println("fi.Size() =>", fi.Size())
	fmt.Println("fi.IsDir() =>", fi.IsDir())
	fmt.Println("fi.ModTime() =>", fi.ModTime())
	fmt.Println("fi.Mode() =>", fi.Mode())
	fmt.Println("\n==================================================================\n")

}
