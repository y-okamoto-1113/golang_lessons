package main

import (
	"fmt"
	"reflect"
)

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
	type Point struct {
		x, y int
	}
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

}
