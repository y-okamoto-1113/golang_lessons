package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	defer f.Close()
	f.WriteString("test")
	fmt.Println("f =>", f)
	fmt.Println("err =>", err)
	fmt.Println("\n==================================================================\n")

	bs := make([]byte, 128)
	fmt.Println("bs =>", bs)
	fmt.Println("\n==================================================================\n")

	// `bs` という変数にファイルの中身を突っ込んでる。
	// Rubyと異なり、ファイルを読み込んだ中身を別の変数に突っ込む。
	file_length, err := f.Read(bs) // 実際に読み込んだバイト数を返す。
	fmt.Println("bs =>", bs)
	fmt.Println("file_length =>", file_length)
	fmt.Println("err =>", err)
	fmt.Println("\n==================================================================\n")

	file_contents := bs[:file_length]
	fmt.Println("file_contents =>", file_contents)
	fmt.Println("string(file_contents) =>", string(file_contents)) // 中身は []byte なので、string型に cast する必要がある。
	fmt.Println("\n==================================================================\n")

}
