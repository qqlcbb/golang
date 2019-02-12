package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" // utf-8, 每个中文三字节
	fmt.Printf(s)
	fmt.Println()

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)  ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s)) // 9

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c  ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		//  (0 Y)  (1 e)  (2 s)  (3 我)  (4 爱)  (5 慕)  (6 课)  (7 网)  (8 !)
		fmt.Printf(" (%d %c) ", i, ch)
	}
}
