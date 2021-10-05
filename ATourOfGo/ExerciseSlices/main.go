package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println("dx, dy =>", dx, dy)
	pic := make([][]byte, dy)
	fmt.Println("pic =>", pic)
	for y := range pic {
		// fmt.Println("y =>", y)
		pic[y] = make([]byte, dx)
		// fmt.Println("pic[y] =>", pic[y])
		for x := range pic[y] {
			// fmt.Println("x =>", x)
			pic[y][x] = byte(x ^ y)
			// fmt.Println("pic[y][x] =>", pic[y][x])
		}
		fmt.Println("pic[y] => ", pic[y])
		fmt.Println("========================")
	}
	return pic
}

func main() {

	pic.Show(Pic)
}
