package main

import (
	"flag"
	"fmt"
)

// `./main -n 999 -m uy4gwecukygcvf -x`の様に実行する
func main() {
	var (
		max int
		msg string
		x   bool
	)

	flag.IntVar(&max, "n", 32, "処理数の最大値") // オプションの型のポインタ、オプションの名前、オプションのデフォルト値、オプションの説明文
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション") // `BoolVar` はそのオプションが渡されたかを返す。
	flag.Parse()

	fmt.Println("処理数の最大値 =", max)
	fmt.Println("処理メッセージ =", msg)
	fmt.Println("拡張オプション =", x)
	fmt.Println("\n==================================================================\n")
}
