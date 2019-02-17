package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	// 定义变量s是一个slice，zero value for slice is nil
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i * 2 + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// 创建slice，设置长度，空间
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Popping from back")
	tail := s2[len(s2) -1]
	s2 = s2[:len(s2) - 1]

	fmt.Println("front", front)
	fmt.Println("back", tail)
	fmt.Println("s2 = ", s2)
}
