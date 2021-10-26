package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("t =>", t)
	fmt.Println("time.Local =>", time.Local)
	fmt.Println("\n==================================================================\n")

	// offsetはUTC時間から何秒進んでいるかを表す。
	tz, offset := t.Zone()
	fmt.Println("tz =>", tz)
	fmt.Println("offset =>", offset)
	// この書き方はできない。t2.Zone()は複数の戻り値を返す。fmt.Printlnでカンマ区切りで複数のものを出力する場合は、各変数は1つのみの戻り値しか対応していない。
	// fmt.Println("t2.Zone() =>", t2.Zone())
	fmt.Println("\n==================================================================\n")

	t2 := time.Date(2021, 10, 25, 17, 56, 22, 999999999, time.Local) // ミリ秒は9桁まで指定できる。
	fmt.Println("t2 =>", t2)
	fmt.Println("t2.Year() =>", t2.Year())
	fmt.Println("t2.YearDay() =>", t2.YearDay())
	fmt.Println("t2.Month() =>", t2.Month())
	fmt.Println("t2.Weekday() =>", t2.Weekday())
	fmt.Println("t2.Day() =>", t2.Day())
	fmt.Println("t2.Hour() =>", t2.Hour())
	fmt.Println("t2.Minute() =>", t2.Minute())
	fmt.Println("t2.Second() =>", t2.Second())
	fmt.Println("t2.Nanosecond() =>", t2.Nanosecond())
	fmt.Println("\n==================================================================\n")

	// 時間の単位
	fmt.Println("time.Hour =>", time.Hour)
	fmt.Println("time.Hour =>", time.Hour)
	fmt.Println("time.Minute =>", time.Minute)
	fmt.Println("time.Second =>", time.Second)
	fmt.Println("time.Millisecond =>", time.Millisecond)
	fmt.Println("time.Microsecond =>", time.Microsecond)
	fmt.Println("time.Nanosecond =>", time.Nanosecond)
	fmt.Println("\n==================================================================\n")

	t3, err := time.ParseDuration("2h30m50s44ms")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("t3 =>", t3)
	fmt.Println("\n==================================================================\n")

	// 時間の追加
	t = t.Add(1*time.Hour + 30*time.Minute)
	fmt.Println("t.Add =>", t)
	fmt.Println("\n==================================================================\n")

	// 時間の差分
	fmt.Println("t2.Sub(t) =>", t2.Sub(t))
	fmt.Println("\n==================================================================\n")

	// `timezone` が異なっていても、時差分だけ埋めた日時が全く同じなら True を返す。
	t5 := time.Date(2015, 10, 10, 0, 0, 0, 0, time.Local)
	t6 := time.Date(2015, 10, 10, 0, 0, 0, 0, time.UTC)
	fmt.Println("t5.Equal(t6) =>", t5.Equal(t6))
	fmt.Println("\n==================================================================\n")

	t7 := time.Date(2015, 10, 10, 00, 00, 00, 00, time.Local)
	t8 := time.Date(2022, 12, 21, 22, 22, 22, 22, time.Local)
	fmt.Println("t7.Before(t8) =>", t7.Before(t8))
	fmt.Println("t7.After(t8) =>", t7.After(t8))
	fmt.Println("\n==================================================================\n")

	// ????????????????????????????????????????????
	// Parseのlayoutが意味わからん。
	layout := "2006/01/02 15:04:05"
	value := "2021/10/26 22:22:22"
	t9, err := time.Parse(layout, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("t9 =>", t9)
	fmt.Println("\n==================================================================\n")

	t10 := time.Date(2021, 12, 21, 22, 22, 22, 22, time.Local)
	fmt.Println("t10.Format(time.RFC3339) =>", t10.Format(time.RFC3339))
	fmt.Println("t10.Format(time.RubyDate) =>", t10.Format(time.RubyDate))
	fmt.Println("\n==================================================================\n")

	fmt.Println("t10.Unix() =>", t10.Unix())
	fmt.Println("t10.UTC() =>", t10.UTC())
	fmt.Println("t10.UTC().Local() =>", t10.UTC().Local())
	fmt.Println("\n==================================================================\n")

	ch := time.Tick(2 * time.Second)
	for {
		t := <-ch
		fmt.Println("reflect.TypeOf(t) =>", reflect.TypeOf(t))
		fmt.Println("t =>", t)
	}
	fmt.Println("\n==================================================================\n")
}
