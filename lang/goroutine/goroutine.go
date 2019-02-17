package main

import (
	"fmt"
	"time"
)

func main() {
	// 开一个a的数组
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("goroutine from %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
