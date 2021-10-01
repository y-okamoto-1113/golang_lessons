package main

import (
	. "fmt"  // パッケージ名を指定せずに関数を使える。
)

// init関数は、main()関数の前に実行される。
// init関数は、他の関数と異なり、重複定義できる。定義した順番で実行される。
func init() {
	Println(1)
}

func init() {
	Println(2)
}

func main() {
	Println(4)
	Println("aaaaaa")
}


func init() {
	Println(3)
}
