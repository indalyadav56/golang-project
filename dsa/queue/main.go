package main

import "fmt"

type Queue struct {
	data []int
	size int
}

func NewQueue(size int) *Queue {
	return &Queue{
		data: make([]int, 0),
		size: size,
	}
}

func (q *Queue) Enqueue(num int) {
	q.data = append(q.data, num)
}

func (q *Queue) Print() {
	for i, val := range q.data {
		fmt.Println(i, val)
	}
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}

	ele := q.data[0]
	q.data = q.data[1:]

	return ele
}

func main() {

	q := NewQueue(3)

	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)

	q.Print()

	fmt.Println("after dequeue")

	q.Dequeue()

	q.Print()

}
