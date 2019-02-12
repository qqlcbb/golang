package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// 从第2个到第6个
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])

	s1 := arr[2:]
	// 经过函数updateSlice，原始数组对应下标的值也会被改变
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	s1 = arr[2:6]
	s2 := s1[3:5] // s1[3], s1[4]
	fmt.Println("gggggg:", s1, s2)

	s3 := append(s2, 10)

	// s4,s5 不再是arr的view，因为超过了原来设定的长度
	s4 := append(s3, 11) // 新的arr
	s5 := append(s4, 12) // 新的arr
	fmt.Println("s3,s4,s4 = ", s3, s4, s5);
	fmt.Println("arr = ", arr);
}
