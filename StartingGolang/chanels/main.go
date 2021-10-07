package main

import (
	"fmt"
)

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go receiver(ch)
	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	fmt.Println("\n==================================================================\n")

	ch2 := make(chan string, 10)
	ch2 <- "Apple"
	fmt.Println("len(ch) =>", len(ch2))
	fmt.Println("cap(ch) =>", cap(ch2))
	ch2 <- "Jonathan Ive"
	ch2 <- "Steve Jobs"
	fmt.Println("len(ch) =>", len(ch2))
	fmt.Println("cap(ch) =>", cap(ch2))
	fmt.Println("\n==================================================================\n")

	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println(<-ch3) // FIFO（先入れ先出し）なので、上から順番に出力される。バッファが空っぽになった chanel からデータを受信しようとするとランタイムエラー
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)

	fmt.Println("\n==================================================================\n")
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	close(ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3) // close した chanel からデータを受信する場合、バッファが空っぽになると型のゼロ値が返り続ける。エラーにならない。
	fmt.Println("\n==================================================================\n")

	ch4 := make(chan int, 3)
	ch4 <- 10
	ch4 <- 20
	ch4 <- 30
	close(ch4)
	i, ok := <-ch4
	fmt.Println("i =>", i, "\nok =>", ok)
	fmt.Println("==========")
	i, ok = <-ch4
	fmt.Println("i =>", i, "\nok =>", ok)
	fmt.Println("==========")
	i, ok = <-ch4
	fmt.Println("i =>", i, "\nok =>", ok)
	fmt.Println("==========")
	i, ok = <-ch4
	fmt.Println("i =>", i, "\nok =>", ok) // バッファが0 && chaenel が close している状態で初めて false が返る。close しててもバッファがある場合は true が返る
	fmt.Println("==========")
	i, ok = <-ch4
	fmt.Println("i =>", i, "\nok =>", ok)
	fmt.Println("\n==================================================================\n")

	// select
	ch5 := make(chan int, 1)
	ch6 := make(chan int, 1)
	ch7 := make(chan int, 1)
	ch5 <- 55
	ch6 <- 66
	// 複数のcaseが成り立つ場合はランダムに実行される。
	select {
	case <-ch5:
		fmt.Println("ch5から受信")
	case <-ch6:
		fmt.Println("ch6から受信")
	case ch7 <- 77:
		fmt.Println("ch7へ送信")
	default:
		fmt.Println("ここへは到達しない。全てのcase節が処理継続不可の場合に実行される。")
	}

}
