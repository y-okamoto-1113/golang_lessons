package main

import (
	"fmt"
	"time"

	greeting "github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do(time.Now()))
}
