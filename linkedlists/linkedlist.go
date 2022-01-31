package linkedlists

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func New(value int) *LinkedList {
	return &LinkedList{
		Value: value,
	}
}

func (l *LinkedList) AddNode(node *LinkedList) *LinkedList {
	if l.Next == nil {
		l.Next = node
		return node
	}
	return l.Next.AddNode(node)
}

func FromArray(arr []int, head *LinkedList) *LinkedList {
	if len(arr) == 0 {
		return head
	}
	return FromArray(arr[1:], head.AddNode(New(arr[0])))
}

func (l *LinkedList) Traverse(callback func(node *LinkedList)) {
	callback(l)
	if l.Next != nil {
		l.Next.Traverse(callback)
	}
}

func (l *LinkedList) Print() {
	l.Traverse(func(node *LinkedList) {
		fmt.Println(node.Value)
	})
}

func (l *LinkedList) ToArray() []int {
	ch := make(chan int)
	arr := []int{}
	go l.Traverse(func(n *LinkedList) {
		ch <- n.Value
		if n.Next == nil {
			close(ch)
		}
	})
	for val := range ch {
		arr = append(arr, val)
	}
	return arr
}
