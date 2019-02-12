package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

// 为结构体定义方法，不是在结构体内定义
func (node Node) Print() {
	fmt.Println(node.Value)
}

// 要使用指针才可以修改
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node, ignored.")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}