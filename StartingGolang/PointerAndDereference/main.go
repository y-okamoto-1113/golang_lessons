package main

import (
	"fmt"
)

func main() {
	var pointer *int
	fmt.Println(pointer)
	fmt.Println("\n==================================================================\n")

	i := 100
	p1 := &i // `&`はアドレス演算子。変数の値ではなく、値が格納されているメモリの住所を返す（ポインタを返す）
	fmt.Println(p1)
	i = 222
	fmt.Println("i =>", i)
	fmt.Println("p1 =>", p1)   // ポインタ（メモリの住所）を返す。
	fmt.Println("*p1 =>", *p1) // メモリの住所（ポインタ）に保存されている実値を返す（デリファレンス）
	fmt.Println("\n==================================================================\n")

	*p1 = 333
	fmt.Println("i =>", i)
	fmt.Println("p1 =>", p1)   // ポインタ（メモリの住所）を返す。
	fmt.Println("*p1 =>", *p1) // メモリの住所（ポインタ）に保存されている実値を返す（デリファレンス）
	fmt.Println("\n==================================================================\n")

	// ポインタを使わなければ、`i`と`p`は別の変数になる（別々のメモリ住所に格納される）。つまり、`i`の値を`p`にコピーして作られる。call by value.
	// しかし、以下では、ポインタを渡して、ポインタからデリファレンスして値を更新しているので、i == p なので、どっちかを更新すると両方更新される。call by refecrence.
	i = 9
	fmt.Println("i =>", i)
	inc(&i)
	fmt.Println("i =>", i)
	inc(&i)
	fmt.Println("i =>", i)
	inc(&i)
	fmt.Println("i =>", i)
	inc(&i)
	fmt.Println("i =>", i)
	fmt.Println("\n==================================================================\n")

	// 配列とポインタ
	a := [5]string{"apple", "beams", "cnet", "domain", "elephant"}
	ap := &a
	fmt.Println(ap)
	fmt.Println("(*ap)[1] =>", (*ap)[1])
	fmt.Println("ap[1] =>", ap[1]) // 本来であれば、`(*ap)[1]`のように ap配列のポインタ（`*ap`）対してインデックス番号1番の値を取得するように記述しないといけないが、Goはコンパイラが自動で検知してくれる。
	fmt.Println("\n==================================================================\n")

	a[1] = "Banana!!!"
	(*ap)[2] = "cucumber??"
	ap[3] = "Dinner With Me --- !?"
	fmt.Println(a)
	fmt.Println(ap)
	fmt.Println("\n==================================================================\n")

	// 文字列とポインタ
	s := ""
	sp := &s
	fmt.Println("s =>", s)
	fmt.Println("&s =>", &s)
	for _, v := range []string{"A", "B", "C", "D", "E", "F"} {
		*sp += v
		fmt.Println("s =>", s)
		fmt.Println("*sp =>", *sp)
		fmt.Println("&s =>", &s)
	}

}

func inc(p *int) {
	*p++ // pをデリファレンスして+1増分
}
