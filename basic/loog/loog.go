package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		panic(fmt.Sprint("Wrong n: %d", n))
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb)  + result
	}
	return result
}

func PrintField(filename string) {
	file , err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		printFileContent(file)
	}
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 省略起始条件，递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func foreach() {
	// 死循环
	for {
		fmt.Println("abd")
	}
}

func main() {
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(13))
	fmt.Println(convertToBin(2))
	PrintField("basic/basic/abc.txt")

	s := `abc"d"
	kkkk
	3333
	`
	printFileContent(strings.NewReader(s));
}
