package main

import (
	"fmt"
	"math/rand"
	"time"
)


func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)

		fmt.Printf("worker %d receiver %d\n", id, n)
	}
}

func createWork(id int) chan <-int{
	c := make(chan int)
	// 创建一个goroutine 对一个chan进行接收
	go worker(id, c)
	// 返回一个chan
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	} ()
	return out

}

func main() {
	var c1, c2 = generator(), generator()
	var workerd = createWork(0)
	// 把收到的信息存放到一个slice里面
	var values []int
	// 函数会返回一个chan，在十秒钟后发送东西给这个chan
	tm := time.After(10 * time.Second)
	// 每隔一秒送一个值过来
	tick := time.Tick(time.Second)
	for {
		// 可以发送数据给nil
		var activeWork chan <-int
		var activeValues int
		// 当检测到收到数据，复制，再发给另外一个worker
		if len(values) > 0 {
			activeWork = workerd
			activeValues = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWork <- activeValues:
			values = values[1:]
		case <- time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <- tick:
			fmt.Println("queue len = ", len(values))
		case <- tm:
			fmt.Println("bye")
			return
		}
	}
}
