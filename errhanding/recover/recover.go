package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Errorm occurred:", err)
		} else {
			panic("I don't know what to do")
		}
	}()
	// panic(errors.New("this is error"))
	// b := 0
	// a := 5/b
	// fmt.Println(a)
	panic("123")
}

func main() {
	tryRecover()
}
