package main

import (
	"fmt"
)

type MyFuncType func()

type MyCustomType []int

func (c MyCustomType) SumOf() {
	for i, v := range c {
		fmt.Println("index :==>", i, v)
	}
}

func main() {
	// a := &MyCustomType{}

	// *a = append(*a, 10, 10, 10)

	// a.SumOf()

	// s1 := [...]int{1, 2, 3, 4, 5}
	// s1 := []int{1, 2, 3, 4, 5}

	// s1 = append(s1, 6, 7, 8)

	// fmt.Println(s1)

	s1 := make([]int, 1)
	s1 = append(s1, 10, 20, 30, 40, 50, 60)
	fmt.Println(len(s1), cap(s1))

	// fmt.Println(s1, len(s1), cap(s1))

	slice := append(s1, 10) // append with existing slices and return new slices with updatec one
	slice[0] = 100

	fmt.Println(s1, slice) // [0] [100 10 20 30]

	// s1 := make([]int, 10, 100)
	// TestSlice(s1)
	// fmt.Println(s1)

	// [0 10 20 30]
	// [100]

}

func TestSlice(nums []int) {
	nums = append(nums, 10, 20, 30)
	nums[0] = 100
	fmt.Println(nums)
}

//
//
//
//

type SomeType struct {
}

// pointer receiver function
func (a *SomeType) Add() {

}

// pointer receiver function
func (a SomeType) Add2() {

}

type MyStruct struct {
	SomeType
}
