package main

import "fmt"

func main() {
	sl := []int{1, 2, 3}

	// sliceheader{
	// 	Data (Addr)
	// 	Len
	// 	Cap
	// }

	var ptr *[]int
	fmt.Println(ptr)

	fmt.Printf("Before modification slice header Addr: %p\n", &sl)
	fmt.Printf("Before modification backing array Addr: %p\n", sl)

	SlicePointer(&sl)

	fmt.Printf("After modification slice header Addr: %p\n", &sl)
	fmt.Printf("After modification backing array Addr: %p\n", sl)
}

func SlicePointer(slPtr *[]int) {
	fmt.Printf("SlicePointer - Addr of slice header param: %p\n", slPtr)
	fmt.Printf("SlicePointer - Addr of backing array before append: %p\n", *slPtr)

	*slPtr = append(*slPtr, 100, 200, 300)

	fmt.Printf("SlicePointer - Addr of backing array after append: %p\n", *slPtr)
}
