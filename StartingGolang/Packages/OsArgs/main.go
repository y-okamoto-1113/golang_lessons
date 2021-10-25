package main

import (
	"fmt"
	"os"
)

// os.Args でコマンドライン引数を取るには、go build した結果に対して実行する必要がある。
// なので、実行するには、`go bild ./main.go` を実行後に `./main 111 222 333 444` のようにする必要がある。
func main() {
	// os.Args は string型 の slice が入ってる。
	fmt.Println("os.Args =>", os.Args)
	fmt.Println("\n==================================================================\n")

	fmt.Println("os.Args[0] =>", os.Args[0])
	fmt.Println("os.Args[1] =>", os.Args[1])
	fmt.Println("os.Args[2] =>", os.Args[2])
	fmt.Println("os.Args[3] =>", os.Args[3])
	fmt.Println("os.Args[4] =>", os.Args[4])
	fmt.Println("\n==================================================================\n")

	fmt.Println("len(os.Args) =>", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}

}
