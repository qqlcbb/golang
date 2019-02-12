package main

import (
	"fmt"
	"test/queue"
)

func main() {

	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push(111)
	fmt.Println(q.Pop())
}
