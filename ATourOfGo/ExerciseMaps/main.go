package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	str := strings.Fields(s)
	// str := strings.Split(s, " ") // Split は空白が複数連続した場合に対応できないので、Fieldsを使う。
	m := make(map[string]int)
	fmt.Println(str)

	for _, v := range str {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
