package main

import "fmt"

// 函数外一定要用var, 所有的变量必须要被用到
var (
	aa = 3
	bb = 4
	cc = 5
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableIntValue() {
	var a int = 3
	var s string = "aaabbb"
	fmt.Println(a, s)
}

// 可以自行识别变量类型
func variableStringValue() {
	var a,b,c = 3, 4, 5
	var s = "ggg"
	fmt.Println(a, b, c, s)
}

// 推荐，冒号定义变量，第二次用不需要冒号，否则重复定义变量，只能在函数里使用
func variableShorterValue() {
	a,b,c := 3, 4, 5
	s := "ggg"
	fmt.Println(a, b, c, s)
}

func consts() {
	const filename = "abc.txt"
	const a, b int = 3, 4
	fmt.Println(filename, a, b)
}

func enms() {
	const(
		cpp = 0
		jave = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, jave, python, golang)
}

func main() {
	fmt.Println("hellow word")
	variableZeroValue()
	variableIntValue()
	variableStringValue()
	variableShorterValue()
	fmt.Println(aa, bb, cc)
	consts()
	enms()
}
