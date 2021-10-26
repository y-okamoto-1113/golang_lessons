package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ディレクトリが存在するとエラーが起きる
	err := os.Mkdir("foo", 0770)
	if err != nil {
		// log.Fatal(err) // この処理だと、ここで処理が `exit` する。今回は `MkdirAll` の処理も見たいので敢えてfmtを使い `exit` しないようにしている。
		fmt.Println(err)
	}

	// `MkdirAll` はディレクトリが存在していてもエラーにならない。
	err = os.MkdirAll("foo/bar/buzz/test/sample/deep/nest", 0775)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
