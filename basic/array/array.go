package main

import "fmt"

func printArr(arr [5]int) {
	for i, v := range(arr) {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i ++ {
		fmt.Println(arr3[i])
	}
	// range 获取数组下标
	for i := range(arr3) {
		fmt.Println(arr3[i])
	}
	// range 获取数组下标和值
	for i, v := range(arr3) {
		fmt.Println(i, v)
	}
	// 只获取值
	for _, v := range(arr3) {
		fmt.Println(v)
	}

	printArr(arr1)
}
