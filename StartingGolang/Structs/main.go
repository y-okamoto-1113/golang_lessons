package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"golang_lessons/StartingGolang/Structs/foo"
)

type Point struct {
	x, y int
}

// メソッド定義
func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.x, p.y)
}

func main() {
	// 複数の方を一気に定義している。
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)

	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Beauty", "Canada", "DT"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)

	fmt.Println("pair =>", pair)
	fmt.Println("reflect.TypeOf(pair) =>", reflect.TypeOf(pair))
	fmt.Println("\n==================================================================\n")
	fmt.Println("strs =>", strs)
	fmt.Println("reflect.TypeOf(strs) =>", reflect.TypeOf(strs))
	fmt.Println("\n==================================================================\n")
	fmt.Println("amap =>", amap)
	fmt.Println("reflect.TypeOf(amap) =>", reflect.TypeOf(amap))
	fmt.Println("\n==================================================================\n")
	fmt.Println("ich =>", ich)
	fmt.Println("reflect.TypeOf(ich) =>", reflect.TypeOf(ich))
	fmt.Println("\n==================================================================\n")

	// struct
	aaa := Point{x: 1, y: 2}
	fmt.Println(aaa)
	fmt.Println(aaa.x)
	fmt.Println("\n==================================================================\n")

	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed // `Feed Feed`と記述したのと同じになる。フィールド名を省略した場合は「フィールド名 == 型名」となる。
	}

	animal := Animal{
		Name: "Rabbit",
		Feed: Feed{Name: "Carrot", Amount: 10},
	}
	fmt.Println("animal =>", animal)
	fmt.Println("animal.Name =>", animal.Name)
	fmt.Println("animal.Feed.Name =>", animal.Feed.Name)
	fmt.Println("animal.Feed.Amount =>", animal.Feed.Amount)
	fmt.Println("animal.Amount =>", animal.Amount) // フィールド名が一意に定まる場合は、中間のフィールド名を省略できる。
	fmt.Println("\n==================================================================\n")

	type Base struct {
		Id    uint
		Owner string
	}

	type A struct {
		Base
		Name string
		Area string
	}

	type B struct {
		Base
		Title  string
		Bodies []string
	}

	a := A{
		Base: Base{Id: 1, Owner: "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}

	b := B{
		Base:   Base{Id: 2, Owner: "Hanako"},
		Title:  "no title",
		Bodies: []string{"test", "bodies", "heehaa!"},
	}

	fmt.Println(a.Id)
	fmt.Println(a.Owner)
	fmt.Println(b.Id)
	fmt.Println(b.Owner)
	fmt.Println("\n==================================================================\n")

	// 構造体とポインタ
	ap := &a
	fmt.Println("a =>", a)
	fmt.Println("ap =>", ap)
	fmt.Println("*ap =>", *ap)
	fmt.Println("\n==================================================================\n")
	(*ap).Name = "Jiro" // 演算子の優先順位の関係で、デリファレンスの部分は`()`で囲む。
	fmt.Println("a =>", a)
	fmt.Println("ap =>", ap)
	fmt.Println("*ap =>", *ap)
	fmt.Println("\n==================================================================\n")
	// わざわざ`*`を使ってデリファレンスしなくても更新できる
	ap.Name = "Suzuki Ichiro"
	fmt.Println("a =>", a)
	fmt.Println("ap =>", ap)
	fmt.Println("*ap =>", *ap)
	fmt.Println("\n==================================================================\n")

	p := Point{x: 7, y: 9}
	swap := func(p *Point) {
		p.x = p.x * 100
		p.y = p.y * 100
	}
	swap(&p)
	fmt.Println(p)
	fmt.Println("\n==================================================================\n")

	// メソッド使用
	pp := &Point{x: 5, y: 12}
	pp.Render()
	fmt.Println("\n==================================================================\n")

	// 構造体の公開・非公開
	t := &foo.T{Field1: 999}
	fmt.Println("t =>", t)
	fmt.Println("*t =>", *t)
	t.Field1 = 1234567
	fmt.Println("t.Field1 =>", t.Field1)
	(*t).Field1 = 98765
	fmt.Println("t.Method1() =>", t.Method1())
	// t.field2 // これは非公開フィールドなのでコンパイルエラー
	// t.method2 // これは非公開メソッドなのでコンパイルエラー
	fmt.Println("\n==================================================================\n")

	// タグ
	type User struct {
		Id   uint   "ユーザーID"
		Name string "名前"
		Age  uint   "年齢"
	}
	u := User{Id: 1, Name: "Kato", Age: 39}
	fmt.Println("u =>", u)
	fmt.Println("\n==================================================================\n")

	tt := reflect.TypeOf(u)
	fmt.Println("tt =>", tt)
	fmt.Println("tt.NumField() =>", tt.NumField())
	for i := 0; i < tt.NumField(); i++ {
		f := tt.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
	fmt.Println("\n==================================================================\n")

	type Staff struct {
		Id       uint   `json: "staff_id"`
		Name     string `json: "staff_name"`
		Age      uint   `json: "staff_age"`
		Position string `json: "staff_position"`
	}
	staff := Staff{Id: 77, Name: "John", Age: 21, Position: "CTO"}
	fmt.Println("staff =>", staff)
	bs, _ := json.Marshal(staff)
	fmt.Println("reflect.TypeOf(bs) =>", reflect.TypeOf(bs))
	fmt.Println("string(bs) =>", string(bs))
	fmt.Println("\n==================================================================\n")

}
