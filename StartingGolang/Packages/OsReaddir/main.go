package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	f, err := os.Open(".") // カレントディレクトリを開く
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fis, err := f.Readdir(0) // `Readdir` は `os.File型` のメソッド。メソッド実行に成功した場合は []os.FileInfo型 を返す
	for i, fi := range fis {
		fmt.Println("\n==================================================================\n")
		if fi.IsDir() {
			fmt.Println("これはディレクトリ")
		}else{
			fmt.Println("これはファイル")
		}
		fmt.Println(i, fi)
		fmt.Println("reflect.TypeOf(fi) =>", reflect.TypeOf(fi))
	}

}
