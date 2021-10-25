package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("test.txt")
	fmt.Println("f =>", f)
	fmt.Println("\n==================================================================\n")

	f.Write([]byte("Hello, World!!!\n\n\ntest"))
	f.WriteAt([]byte("Golang"), 0)  // 0番目の文字（オフセット位置）から `Golang` に書き換えている。
	fmt.Println("\n==================================================================\n")

	f.WriteString("111111111 harahetta 111111111")
	fmt.Println("\n==================================================================\n")

	f.Seek(0, os.SEEK_END)
	f.WriteString("222222")
	fmt.Println("\n==================================================================\n")

}
