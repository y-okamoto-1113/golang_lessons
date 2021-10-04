package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	values := [...]float64{2, 3, 4, 5, 6, 7, 8, 9, 10} // float64にしても数値リテラルで無理やりIntegerに型推論されるので、forでそれぞれfloat64に型変換する必要がある。
	for _, i := range values {
		x := float64(i) // iはなぜかループ処理で使っていて上書きできないので別の変数を定義する必要がある。
		fmt.Println(x, "の平方根: ", Sqrt(x))
		fmt.Println(x, "の平方根: ", math.Sqrt(x))
	}
}
