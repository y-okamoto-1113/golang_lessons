package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	fmt.Println("file =>", file)
	fmt.Println("err =>", err)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 1024)
	fmt.Println("buff 111=>", buff)
	count, err := file.Read(buff) // byte型の変数を定義して、そこにデータを突っ込んでる。byte型で定義した capacity 以上は代入できない。 countは文字数を返す。
	fmt.Println("buff 222=>", buff)
	fmt.Println("count =>", count)
	fmt.Println("err =>", err)
	fmt.Println("buff[:count]=>", buff[:count])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buff))

}
