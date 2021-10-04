//! これは無限ループ。平行処理の処理内容確認のために実装した。
package main

import (
	"fmt"
)

func sub() {
	for {
		fmt.Println("sub loop")
	}
}

func main() {
	go sub()
	for {
		fmt.Println("main loop")
	}
}
