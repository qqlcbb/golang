package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+" :
		return a + b, nil
	case "-" :
		return a - b, nil
	case "/" :
		q, _ := div(a, b)
		return q, nil
	case "*" :
		return a * b, nil
	default :
		return 0, fmt.Errorf("unsupported operation:" + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a * b
}

// 使用函数作为参数
func apply(op func (int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function %s with args" + "(%d, &d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(eval(4, 2 , "g"));
	q, r := div (4, 2)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3,4))
}
