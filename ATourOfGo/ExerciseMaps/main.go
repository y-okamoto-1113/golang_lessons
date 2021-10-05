package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	str := strings.Split(s, " ")
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
