package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	fmt.Printf("GAME START\n")
	firstWord := "り"
	start(firstWord)

}

func start(word string){
	fmt.Printf("最初の文字は: %s\n", word)

	text := input()
	fmt.Printf(text)
}

func input() (string){
	fmt.Printf("文字を入力してください\n")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

