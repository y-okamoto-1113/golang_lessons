package main

import (
	"fmt"
	"reflect"
)

// fibonacci関数 は 引数を取らない。
// 返り値は、function（無名関数）
// 返ってくる function（無名関数）の返り値は int。
func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// fibonacci関数を実行したので、無名関数が返ってきてる。
	f := fibonacci()
	fmt.Println("reflect.TypeOf(f) =>", reflect.TypeOf(f))
	for i := 0; i < 20; i++ {
		// 無名関数fに`()`がついてるので、関数を実行している。
		fmt.Println(f())
	}
}
