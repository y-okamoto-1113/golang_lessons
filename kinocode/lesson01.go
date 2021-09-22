package main
import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Good Morning!")
	fmt.Println("Good Afternoon!")
	fmt.Println("Good Evening!")

	// 変数名を再度宣言することはできないので、変数名を分けている
	var num1 int = 1
	fmt.Println(num1)
	var num2 = 2
	fmt.Println(num2)
	num3 := 3
	fmt.Println(num3)

	// case sensitive (大文字・小文字は区別される)
	NUM := 4
	fmt.Println(NUM)
	Num := 5
	fmt.Println(Num)

	// Integer
	// https://blog.y-yuki.net/entry/2017/04/27/000000
	fmt.Println("===================================================================")
	fmt.Println("数値型について")
	fmt.Println("===================================================================")
	// int8 は 8bit なので、0~255までの値を取るが、これは -128 ~ 127 の値を取る。
	// 1bit目は、プラス・マイナス符号になる。
	// -128, +127を超えると「overflows int8」となる
	fmt.Println("========== int8について ==========")
	var int8_num int8 = 127
	fmt.Println(int8_num)
	fmt.Println(reflect.TypeOf(int8_num))

	// uint8は 8bit でプラスの値のみ。 0~255までの値を取る。
	fmt.Println("========== uint8について ==========")
	var uint8_num uint8 = 255
	fmt.Println(uint8_num)
	fmt.Println(reflect.TypeOf(uint8_num))

	// byte は uint8 のエイリアス
	fmt.Println("========== byteについて ==========")
	var byte_num byte = 255
	fmt.Println(byte_num)
	fmt.Println(reflect.TypeOf(byte_num))

	// int16 は 16bit -32768 ~ +32767 の値を取る。
	fmt.Println("========== int16について ==========")
	var int16_num int16 = 32767
	fmt.Println(int16_num)
	fmt.Println(reflect.TypeOf(int16_num))

	// int32 は -2147483648 ~ +2147483647 の値を取る。
	fmt.Println("========== int32について ==========")
	var int32_num int32 = 2147483647
	fmt.Println(int32_num)
	fmt.Println(reflect.TypeOf(int32_num))

	// int64 は -9223372036854775808 ~ +9223372036854775807 の値を取る。
	fmt.Println("========== int64について ==========")
	var int64_num int64 = 9223372036854775807
	fmt.Println(int64_num)
	fmt.Println(reflect.TypeOf(int64_num))

	// Float
	// 整数と小数点で合計8桁目までしか表示しない
	fmt.Println("========== float32について ==========")
	var float32_num float32 = 1.234567890543298765432456789087654356789864568
	fmt.Println(float32_num)

	// 整数と小数点で合計17桁目までしか表示しない
	fmt.Println("========== float64について ==========")
	var float64_num float64 = 1.12346789098321235678909412345678987654321234567
	fmt.Println(float64_num)


	// String
	fmt.Println("===================================================================")
	fmt.Println("文字列型について")
	fmt.Println("===================================================================")
	var string_a string = "Hello World from a!"
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	string_b := "Hello World from b!"
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))


	// Boolean
	fmt.Println("===================================================================")
	fmt.Println("真偽値型について")
	fmt.Println("===================================================================")
	var true_1 bool = true
	fmt.Println(true_1)
	fmt.Println(reflect.TypeOf(true_1))
	var false_1 bool = false
	fmt.Println(false_1)
	fmt.Println(reflect.TypeOf(false_1))

	a := 1
	b := 10
	var num_bool bool = a > b
	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))

	// Array
	fmt.Println("===================================================================")
	fmt.Println("配列 について")
	fmt.Println("===================================================================")
	arr := [5]string{
		"okamoto",
		"kunisawa",
		"seki",
		"nishioka",
		"tanaka",
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}


	// 条件分岐
	fmt.Println("===================================================================")
	fmt.Println("条件分岐 について")
	fmt.Println("===================================================================")
	age := 19
	if age >= 20 {
		fmt.Println("adult")
	} else if age >= 6 {
		fmt.Println("child")
	} else {
		fmt.Println("baby")
	}

	apple := 200
	banana := 150
	if sum := apple + banana; sum >= 400 {
		fmt.Println("高過ぎぃぃぃ！！！")
	} else if sum >= 200 {
		fmt.Println("神の見えざる手 が働いていますね！")
	} else {
		fmt.Println("安過ぎぃぃぃ！！！！")
	}



	// ループ処理
	fmt.Println("===================================================================")
	fmt.Println("ループ処理 について")
	fmt.Println("===================================================================")
	for i := 0; i < 12; i++ {
		if i%4 == 0 { continue }
		if i == 11 {break}
		fmt.Println(i)
	}


	// 関数
	fmt.Println("===================================================================")
	fmt.Println("関数作成 について")
	fmt.Println("===================================================================")
	sayHello("Good Morning with Interpolation")
	cal1 := func(x, y, z int) (int, int, int, int){
		addition := x + y + z
		substraction := x - y - z
		division := x / y / z
		multiplication := x * y * z
		return addition, substraction, division, multiplication
	}
	cal2 := func(x, y, z int) (addition, substraction, division, multiplication int){
		addition = x + y + z
		substraction = x - y - z
		division = x / y / z
		multiplication = x * y * z
		return
	}
	fmt.Println( cal1(10,10,10) )
	fmt.Println( cal2(10,10,10) )
	func(greeting string){
		fmt.Println(greeting)
	}("Good Night!")

	lesson1_cal := func(x, y int)(result int){
		result = x + y
		return
	}
	fmt.Println( lesson1_cal(10, 5) )

}

// 関数
func sayHello(greeting string){
	fmt.Println(greeting)
}



