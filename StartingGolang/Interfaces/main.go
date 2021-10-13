package main

import (
	"fmt"
	"reflect"
)

type MyError struct {
	Message string
	ErrCode int
}

// eという変数に、Error()というメソッドを定義
func (e *MyError) Error() string {
	return e.Message
}

// MyError を return しているが、戻り値の型が error型 なのでキャストされてる
func RaiseError() error {
	return &MyError{Message: "1: エラーが発生しました。", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println("err =>", err)
	fmt.Println("reflect.TypeOf(err) =>", reflect.TypeOf(err))
	fmt.Println("err.Error() =>", err.Error())
	// fmt.Println("err.Error() =>", err.Message()) // error型にキャストされてて、MyError型じゃなくなってるので、Message フィールドが存在しない。
	// fmt.Println("err.Error() =>", err.ErrCode()) // error型にキャストされてて、MyError型じゃなくなってるので、ErrCode フィールドが存在しない。

	// 本来の型（MyError型）を使うには、型アサーションする
	fmt.Println("err.(*MyError) =>", err.(*MyError))
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
	fmt.Println("\n==================================================================\n")

	err2 := MyError{Message: "2: エラーが発生しました。", ErrCode: 1234}
	fmt.Println("err2 =>", err2)
	fmt.Println("reflect.TypeOf(err2) =>", reflect.TypeOf(err2))
	fmt.Println("err2.Message =>", err2.Message)
	fmt.Println("err2.ErrCode =>", err2.ErrCode)
	fmt.Println("\n==================================================================\n")

	err3 := &MyError{Message: "3: エラーが発生しました。", ErrCode: 1234}
	fmt.Println("err3 =>", err3)
	fmt.Println("reflect.TypeOf(err3) =>", reflect.TypeOf(err3))
	fmt.Println("err3.Message =>", err3.Message)
	fmt.Println("err3.ErrCode =>", err3.ErrCode)
	fmt.Println("\n==================================================================\n")
}
