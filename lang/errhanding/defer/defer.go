package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"test/lang/errhandingnding/fib"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	// 自己建立错误
	err = errors.New("this is a custom error")
	if (err != nil) {
		// panic(err)
		// fmt.Println("Error:", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(write, f())
	}


}

func main() {
	writeFile("gg.txt")
}
