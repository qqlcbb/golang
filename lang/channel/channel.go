package main

import (
	"fmt"
	"time"
)

// chan作为返回值
func createWork(id int) chan<- int{
	c := make(chan int)
	// 创建一个goroutine 对一个chan进行接收
	go work(id, c)
	// 返回一个chan
	return c
}


func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		// 返回一个chan
		channels[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		// 给channel发送消息
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		// 给channel发送消息
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func work(id int, c chan int) {
	for n := range c {
		// n, ok := <- c
		// if (!ok) {
		// 	break;
		// }
		fmt.Printf("worker %d receiver %d\n", id, n)
	}
}

func bufferChan() {
	// 加入一个缓存区，大小为3, 提升性能
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func choseChan() {
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// 告诉接收方，发送完毕
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	// chanDemo()
	// bufferChan()
	choseChan()
}
