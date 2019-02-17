package adder

import "fmt"

func adder() func(int) int {
	// 自由变量
	sum := 0;
	// 返回一个闭包
	return func(v int) int {
		sum += v
		return sum
	}
}

// 正统函数式编程
type iAdder func(int) (int, iAdder)
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	// 函数体包含（局部变量，自由变量）
	a := adder()
	// 把所有的i相加
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
	fmt.Println()
	// 使用正统函数式编程
	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}

}
