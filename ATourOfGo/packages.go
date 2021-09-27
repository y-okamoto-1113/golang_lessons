package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// https://makiuchi-d.github.io/2017/09/09/qiita-9c4af327bc8502cdcdce.ja.html
	// https://www.spinute.org/go-by-example/random-numbers.html
	fmt.Println(rand.Seed)
	// rand.Intnは疑似乱数を返す。この場合「0~100」の値が固定されている。
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))

	// 疑似乱数はシードに基づいて算出されるが、シード値は1で固定されている。毎回異なる乱数を出力するには、rand.Seedでシード値を更新する必要がある。
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println(rand.Seed)

	fmt.Println(math.Pi)

	aaa := func(x, y string) (string, string) {
		return y, x
	}
	fmt.Println(aaa("word-1111", "word2"))

}
