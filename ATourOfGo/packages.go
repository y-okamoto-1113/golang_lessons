package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/cmplx"
	"math/rand"
	"reflect"
	"strconv"
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
	fmt.Println("\n==================================================================\n")

	aaa := func(x, y string) (string, string) {
		return y, x
	}
	fmt.Println(aaa("word-1111", "word2"))
	fmt.Println("\n==================================================================\n")

	// ちゃんとした乱数生成
	fmt.Println(math.MaxInt64)
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	fmt.Println(rand.Int63())

	// 変数宣言はimportステートメントと同じくまとめて宣言（factored statement）できる
	var (
		boolean       bool       = false
		MaxUint       uint64     = 1<<64 - 1            // ビット演算。2進数で0を64個追加している。
		MaxInt        int64      = 1<<63 - 1            // intはuintと異なり、±符号で1bit使用するので純粋に数値として扱える数値は63bitのみ。
		ComplexNumber complex128 = cmplx.Sqrt(-5 + 12i) // 複素数
	)
	fmt.Printf("boolean変数のType: %T Value: %v\n", boolean, boolean)
	fmt.Printf("MaxUint変数のType: %T Value: %v\n", MaxUint, MaxUint)
	fmt.Printf("MaxInt変数のType:  %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("ComplexNumber変数のType: %T Value: %v\n", ComplexNumber, ComplexNumber)

	fmt.Printf("MaxUint変数と math.MaxUint64 は同じ値？ => %v\n", MaxUint == math.MaxUint64)
	fmt.Printf("MaxInt変数と math.MaxInt64 は同じ値？   => %v\n", MaxInt == math.MaxInt64)
	fmt.Println("\n==================================================================\n")

	// コードポイントについて
	// fmt.Printlnで出力される文字はUTF-8にエンコードされる。UTF-8はコードポイントを16進数表記なので、16真数に変換してUnicodeの頭文字をつける
	fmt.Println("銀")
	fmt.Println('銀') // コードポイントが10進数表記で出力される。

	// コードポイントを整数として取得
	var codepoint int = '銀' // 型がruneの場合 strconv.Itoa で変換できないので敢えて int にしている。
	fmt.Printf("変数codepoint の型: %s \n", reflect.TypeOf(codepoint))
	fmt.Printf("変数codepoint の値: %d \n", codepoint)

	// 16進数に変換
	var CodepointEncoded16Bit string = fmt.Sprintf("%x", codepoint) // 16進数にエンコードするので
	fmt.Printf("変数 CodepointEncoded16Bit の型: %s \n", reflect.TypeOf(CodepointEncoded16Bit))
	fmt.Printf("変数 CodepointEncoded16Bit の値: %s \n", CodepointEncoded16Bit)
	fmt.Println("\u9280")
	fmt.Println(string([]rune{0x9280}))

	// stringメソッドはstring型に型変換してくれるが、数値はコードポイントとして認識される。
	fmt.Println(string("37504"))
	fmt.Println(string("\u9280"))
	fmt.Println("\n==================================================================\n")

	// 型変換
	i := 27515
	fmt.Printf("変数iの型: %T \n", i)
	fmt.Printf("変数iの値: %v \n", i)
	f := float64(i)
	fmt.Printf("変数fの型: %T \n", f)
	fmt.Printf("変数fの値: %v \n", f)
	u := uint(f)
	fmt.Printf("変数uの型: %T \n", u)
	fmt.Printf("変数uの値: %v \n", u)
	// Integer を string に型変換すると、その数値はコードポイントとして認識されるので、strconv
	s := strconv.Itoa(i)
	fmt.Printf("変数sの型: %T \n", s)
	fmt.Printf("変数sの値: %v \n", s)
	s_codepoint := string(i)
	fmt.Printf("変数s_codepointの型: %T \n", s_codepoint)
	fmt.Printf("変数s_codepointの値: %v \n", s_codepoint)
	fmt.Println("\n==================================================================\n")

	// slice
	a := []int{10, 20}
	b := append(a, 30)
	a[0] = 100
	c := append(b, 40)
	b[1] = 999
	d := append(c, 50)
	c[3] = 200
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d) // キャパシティオーバーでポインタ値が変わるので、別オブジェクトになる。
	fmt.Println("\n==================================================================\n")

}
