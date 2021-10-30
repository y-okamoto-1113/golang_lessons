package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(256)
	fmt.Println("乱数生成シード値 256 の場合の rand.Float64() =>", rand.Float64())
	fmt.Println("乱数生成シード値 256 の場合の rand.Float64() =>", rand.Float64())
	fmt.Println("乱数生成シード値 256 の場合の rand.Float64() =>", rand.Float64())
	fmt.Println("\n==================================================================\n")

	// 乱数シード値が毎回変わるので結果も変わる。
	rand.Seed(time.Now().UnixNano())
	fmt.Println("rand.Float64() =>", rand.Float64())
	fmt.Println("rand.Intn(100) =>", rand.Intn(100)) // 0以上100未満の値をランダムに返す。
	fmt.Println("rand.Int31() =>", rand.Int31())
	fmt.Println("rand.Int63() =>", rand.Int63())
	fmt.Println("rand.Uint32() =>", rand.Uint32())
	fmt.Println("\n==================================================================\n")

	// `rand.Seed` のシード値はランタイム全体で共有されるので、シード値を変える場合は都度、新しい擬似乱数生成器を作成する必要がある。
	src := rand.NewSource(time.Now().UnixNano()) // rand.Source型で、新しいシード値を生成する。
	rnd := rand.New(src)                         // シード値を渡して、新しい疑似乱数生成器を作成する。
	fmt.Println("rnd.Float64() =>", rnd.Float64())

}
