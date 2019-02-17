package main

import (
	"fmt"
	"sync"
)

// chan作为返回值
func createWork(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	// 创建一个goroutine 对一个chan进行接收
	go doWork(id, w)
	// 返回一个chan
	return w
}


func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker

	for i := 0; i < 10; i++ {
		// 返回一个chan
		workers[i] = createWork(i, &wg)
	}
	// 等待20个任务完成
	wg.Add(20)

	for i, worker := range workers {
		// 给channel发送消息
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		// 给channel发送消息
		worker.in <- 'A' + i
	}

	wg.Wait()
}

// 定义一个结构
type worker struct {
	in chan int
	done func()
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d receiver %c\n", id, n)
		// 使用另外一个chan来通知事情已经做完
		w.done()
	}
}

func main() {
	chanDemo()
}
