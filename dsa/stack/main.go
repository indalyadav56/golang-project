package main

import "fmt"

type Stack struct {
	data []int
	size int
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]int, 0),
		size: size,
	}
}

func (s *Stack) Push(num int) int {
	if len(s.data) <= s.size {
		s.data = append(s.data, num)
		return num
	}
	return -1
}

func (s *Stack) Print() {
	for i, val := range s.data {
		fmt.Println(i, val)
	}
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}

	lastEle := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return lastEle
}

func main() {
	stack := NewStack(3)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Print()

	stack.Pop()

	fmt.Println("after pop")

	stack.Print()

}
