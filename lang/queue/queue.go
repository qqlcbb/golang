package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	// 函数里限定类型
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}