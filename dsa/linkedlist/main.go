package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data int) {
	node := NewNode(data)

	if l.head == nil {
		l.head = node
		return
	}

	current := l.head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = node

}

func (l *LinkedList) Print() {

	current := l.head
	for current.Next != nil {
		fmt.Println("data:->", l.head.Next.Data)
		current = current.Next
	}

}

func main() {

	ll := NewLinkedList()

	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)
	ll.Insert(50)

	ll.Print()
}
