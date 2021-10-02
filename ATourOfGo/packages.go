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
	"strings"
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
	// 整数の最大値は以下でビット演算で定義されている。
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.1:src/math/const.go;l=39
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
	slice_a := []int{10, 20}
	slice_b := append(slice_a, 30)
	slice_a[0] = 100
	slice_c := append(slice_b, 40)
	slice_b[1] = 999
	slice_d := append(slice_c, 50)
	slice_c[3] = 200
	fmt.Println(slice_a)
	fmt.Println(slice_b)
	fmt.Println(slice_c)
	fmt.Println(slice_d) // キャパシティオーバーでポインタ値が変わるので、別オブジェクトになる。

	// sliceデータは、比較できない。ポインタの値も含む為。
	// slice_e := []int{10, 20, 30}
	// slice_f := []int{10, 20, 30}
	// slice_e == slice_f
	fmt.Println("\n==================================================================\n")

	// array
	ns := [...]int{10, 20, 30}
	ms := [...]int{10, 20, 30}
	fmt.Println(ns == ms)
	fmt.Println("\n==================================================================\n")

	// オーバーフロー
	// 数値リテラルはint型と推論される。uintは型推論されないので、使いたい場合は明記するしかない。
	// res := math.MaxUint64   =>  正の整数リテラルは全てint型と推論される。
	var res uint64 = math.MaxUint64
	fmt.Println(res)
	var uint_1 uint64 = math.MaxUint32
	uint_2 := math.MaxUint32
	uint_3 := uint64(uint_2)
	fmt.Printf("変数uint_1の型: %T, 値: %d \n", uint_1, uint_1)
	fmt.Printf("変数uint_2の型: %T, 値: %d \n", uint_2, uint_2)
	fmt.Printf("変数uint_3の型: %T, 値: %d \n", uint_3, uint_3)
	fmt.Println("\n==================================================================\n")

	// ラップアラウンド
	before := byte(255)
	after := before + 1
	fmt.Println(before) // 255は「0 1111 1111」なので、255と表現される。
	fmt.Println(after)  // 256は「1 0000 0000」であり、byte型は8bitまでの値を見る。このような繰り上がりは無視されて8bitまでの値で計算が行われる。
	fmt.Println("\n==================================================================\n")

	// 浮動小数点数の無限値・非数
	zero := 0.0
	pf := 1.0
	nf := -1.0
	fmt.Println(pf / zero)
	fmt.Println(nf / zero)
	fmt.Println(zero / zero)
	fmt.Println("\n==================================================================\n")

	// iota
	const (
		const_a = iota
		const_b
		const_c
		const_d = 10 + iota
		const_e
		const_f
		const_g = 999
		const_h
		const_i
		const_j = iota // iotaはconstブロック内で定数が定義されるたびに1ずつ増える。ここは10番目なので「9」になる。
		const_k
	)
	fmt.Println(const_a, const_b, const_c, const_d, const_e, const_f, const_g, const_h, const_i, const_j, const_k)

	// 識別子の命名規則
	const (
		朝の挨拶 = "Good Morning!!!"
		昼の挨拶 = "Good Afternoon!!!"
		晩の挨拶 = "Good Evening!!!"
	)
	fmt.Println(朝の挨拶, 昼の挨拶, 晩の挨拶)
	fmt.Println("\n==================================================================\n")

	// slice の length と capacity
	sl := []struct{ Age int }{
		{Age: 10},
		{Age: 20},
		{Age: 30},
		{Age: 40},
		{Age: 50},
	}
	fmt.Println(sl)
	fmt.Printf("len(sl) => %v\n", len(sl))
	fmt.Printf("cap(sl) => %v\n", cap(sl))
	fmt.Println("\n==================================================================\n")

	// slice range
	sl2 := []struct{ height float64 }{
		{height: 180.0},
		{height: 180.1},
		{height: 180.2},
		{height: 175.9},
		{height: 168.5},
	}
	// range 〇〇 で〇〇の長さ0から最後までのインデックスを取得して、順番に実行する。
	// rubyの each を実行したいときに 古典的なfor を書かなくても楽にできる。
	for i := range sl2 {
		fmt.Printf("sl2[%d] => %v \n", i, sl2[i].height)
	}

	// sl2の値を1個1個取得し、vという変数にコピーしてるから、sl2の値は更新されない。
	for _, v := range sl2 {
		v.height *= 10
	}
	fmt.Println(sl2)

	// sl2 を直接参照しているので、値は更新される。
	for i := range sl2 {
		sl2[i].height *= 1000
	}
	fmt.Println(sl2)
	fmt.Println("\n==================================================================\n")

	// スライスのキャパシティ指定。
	var longArray [1_000_000]int
	laSlice1 := longArray[:10]
	fmt.Printf("len(laSlice1) => %v \n", len(laSlice1))
	fmt.Printf("cap(laSlice1) => %v \n", cap(laSlice1))
	// 配列からスライスを作成する際にキャパシティを指定しなければ、コピー元の配列と同じキャパシティになる。
	laSlice2 := longArray[:10:11]
	fmt.Printf("len(laSlice2) => %v \n", len(laSlice2))
	fmt.Printf("cap(laSlice2) => %v \n", cap(laSlice2))
	fmt.Println("\n==================================================================\n")

	// map
	m := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	}
	fmt.Println(m)
	fmt.Println(m["x"])
	// キーの存在チェック
	map_z, ok := m["z"] // 存在すれば value と true が返る。
	println(map_z, ok)
	// キーの削除
	delete(m, "z")
	map_z, ok = m["z"] // 存在しなければ ゼロ値 と false が返る。
	println(map_z, ok)
	// キーの追加
	m["z"] = 99
	map_z, ok = m["z"]
	println(map_z, ok)
	fmt.Println(m)
	fmt.Println("\n==================================================================\n")

	// Split でカウント
	str := "dog dog dog cat dog rabbit cat rabbit"
	split := strings.Split(str, " ")
	fmt.Println(len(split))
	fmt.Println(cap(split))
	fmt.Printf("Type => %T \n", split)
	fmt.Printf("Value => %v \n", split)
	// 要素ごとのカウント
	counts := map[string]int{}
	for _, s := range strings.Split(str, " ") {
		fmt.Println(s)
		counts[s]++
	}
	fmt.Println(counts)
	// mapを使って Boolean のゼロ値を活用
	flags := map[string]bool{}
	flags["dog"] = true
	if flags["dog"] {
		fmt.Println("dog OK!!!!")
	}
	// キャットの値は存在しないのでここで初めて宣言する。値が指定されていない場合は、ゼロ値で初期化される。boolean の場合は false
	if flags["cat"] {
		fmt.Println("cat OK!!!!")
	}
	fmt.Println("\n==================================================================\n")

	// ネストした配列
	// golangでは配列の要素も`{}`で記述する。`[]`はNG！
	nestedArray := [][][]int{
		{
			{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
		},
		{
			{11, 12, 13}, {14, 15, 16}, {17, 18, 19},
		},
		{
			{21, 22, 23}, {24, 25, 26}, {27, 28, 29},
		},
	}
	fmt.Println(nestedArray)
	for i := range nestedArray {
		fmt.Println(nestedArray[i])
	}
	fmt.Println("\n==================================================================\n")

	// interface{}型
	anything := func(a interface{}) {
		fmt.Printf("[value => %v], [type => %T]\n", a, a)
	}
	anything(999)
	anything(3.14)
	anything(4 + 5i)
	anything('忍')
	anything("忍者")
	anything([]int{1, 2, 3, 4, 5})
	fmt.Println("\n==================================================================\n")

	// interface{}型の型変換
	var randVal interface{} = 3.14
	i, isInt := randVal.(int) // int64などは指定できない。
	f, isFloat := randVal.(float64)
	s, isString := randVal.(string)
	r, isRune := randVal.(rune)
	fmt.Printf("i, isInt => %v, %v\n", i, isInt)
	fmt.Printf("f, isFloat => %v, %v\n", f, isFloat)
	fmt.Printf("s, isString => %v, %v\n", s, isString)
	fmt.Printf("r, isRune => %v, %v\n", r, isRune)
	fmt.Println("\n==================================================================\n")

}
