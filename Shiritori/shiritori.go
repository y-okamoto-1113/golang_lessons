package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("GAME START\n")
	firstWord := "り"
	start(firstWord)

}

func start(word string) {
	fmt.Printf("最初の文字は: %s\n", word)
	text := input()
	fmt.Printf("入力された文字は: %s", text)
	judge(text, word)
}

// 標準入力を受け取る
func input() string {
	fmt.Printf("文字を入力してください\n")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

func judge(text, word string) {
	textRune := []rune(text)
	firstWord := string(textRune[0])
	size := len(textRune)

	if firstWord != word {
		fmt.Println("再世の文字が違います。")
		fmt.Printf("入力した最初の文字: %s\n", firstWord)
		fmt.Printf("もう一度入力してください。")
		start(word)
	}

	whiteList := []string{
		"あ", "い", "う", "え", "お",
		"か", "き", "く", "け", "こ",
		"さ", "し", "す", "せ", "そ",
		"た", "ち", "つ", "て", "と",
		"な", "に", "ぬ", "ね", "の",
		"は", "ひ", "ふ", "へ", "ほ",
		"ま", "み", "む", "め", "も",
		"や", "ゆ", "よ",
		"ら", "り", "る", "れ", "ろ",
		"わ", "ゐ", "ゑ", "を", "ん",

		"が", "ぎ", "ぐ", "げ", "ご",
		"ざ", "じ", "ず", "ぜ", "ぞ",
		"だ", "ぢ", "づ", "で", "ど",
		"ば", "び", "ぶ", "べ", "ぼ",
		"ぱ", "ぴ", "ぷ", "ぺ", "ぽ",

		"ぁ", "ぃ", "ぅ", "ぇ", "ぉ",
		"っ",
		"ゃ", "ゅ", "ょ",

		"ア", "イ", "ウ", "エ", "オ",
		"カ", "キ", "ク", "ケ", "コ",
		"サ", "シ", "ス", "セ", "ソ",
		"タ", "チ", "ツ", "テ", "ト",
		"ナ", "ニ", "ヌ", "ネ", "ノ",
		"ハ", "ヒ", "フ", "ヘ", "ホ",
		"マ", "ミ", "ム", "メ", "モ",
		"ヤ", "ユ", "ヨ",
		"ラ", "リ", "ル", "レ", "ロ",
		"ワ", "ヰ", "ヱ", "ヲ",
		"ン",
		"ガ", "ギ", "グ", "ゲ", "ゴ",
		"ザ", "ジ", "ズ", "ゼ", "ゾ",
		"ダ", "ヂ", "ズ", "デ", "ド",
		"バ", "ビ", "ブ", "ベ", "ボ",
		"パ", "ピ", "プ", "ペ", "ポ",

		"ァ", "ィ", "ゥ", "ェ", "ォ",
		"ッ",
		"ャ", "ュ", "ョ",
	}

	for key, value := range text {
		if !inArray(whiteList, value) {
			fmt.Printf(key)
		}
	}

	lastWord := string(textRune[size-1])
	fmt.Printf("最後の文字: %s", lastWord)

	if lastWord == "ん" {
		fmt.Println("最後の文字が「ん」です。")
		end()
	}

	start(lastWord)
}

func inArray(whiteList []string, word string) bool {
	for _, value := range whiteList {
		if word != value {
			return false
		}
		return true
	}
}

func end() {
	fmt.Println("GAME OVER!!!")
	os.Exit(0)
}
