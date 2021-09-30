package main

import (
	"fmt"

	"golang_lessons/StartingGolang/zoo/animals" // 相対パスはコンパイルエラーになるので、GOPATHからの絶対パスを記述
)

func main() {
	fmt.Println("first message!")
	fmt.Println(AppName())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
