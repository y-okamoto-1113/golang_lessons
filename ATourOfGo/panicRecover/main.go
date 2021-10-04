package main

import (
	"fmt"
	"log"
)

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch src.(type) {
			case int:
				log.Println("panic: int=", x)
			case float64:
				log.Println("panic: float64=", x)
			case string:
				log.Println("panic: string=", x)
			default:
				log.Println("panic: unknown value!")
			}
		}
	}()
	panic(src)
}

func main() {
	// panic && recover
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic("panic panic!!!")
	fmt.Println("実行されない！")

	// 上をコメントアウトして以下をコメントオフ
	// testRecover(1111)
}
