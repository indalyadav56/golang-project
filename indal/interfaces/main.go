package main

import "fmt"

func main() {

}

type Interface1 interface {
	Method1(num int) error
}

func NewMethod(i Interface1) {
	i.Method1(10)
	fmt.Println("indal")
}

type Interface2 interface {
	Method2(num int) error
}

type NewStruct struct {
}

func (n *NewStruct) Method1(num int) error {
	fmt.Println("num:-=-=>", 10)
	return nil
}
