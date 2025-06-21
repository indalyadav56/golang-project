package main

import "fmt"

func main() {

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
