package main

import (
	"fmt"
)

func main() {

	// data := 9

	// fmt.Printf("%d, %b \n", data, data)

	// // pointer to struct
	// ptrStruct := PointerStuct{
	// 	Name:  "INDAL",
	// 	Email: "indal@gmail.com",
	// }

	// fmt.Println(ptrStruct)

	// ptr := &ptrStruct

	// ptr.Name = "newName"
	// ptr.Email = "newEmail"

	// fmt.Println(*ptr)

	// // pointer to slices
	// slices := []int{1, 2, 3, 4, 5}

	// sp := &slices

	// slices = append(*sp, 6, 7, 8)

	// fmt.Println(slices)
	// fmt.Println(*sp)

	var p *int

	// * opeartion (defaine o) and dereferenceing

	// pointer type(base type should be only int)

	// str := "indal"
	x := 10

	p = &x

	*p = 1000

	fmt.Println(*p)
	// p =

	// x := 10
	// p = &x

	// fmt.Println(*p)

}

type PointerStuct struct {
	Name  string
	Email string
}
